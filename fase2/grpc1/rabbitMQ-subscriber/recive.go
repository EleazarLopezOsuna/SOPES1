package main

import (
	"context"
	encoder "encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
	amqp "github.com/rabbitmq/amqp091-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type getRabbitData struct { // struct para obtener data de RabbitMQ
	Game_Id   string `json:"game_id"`
	Players   string `json:"players"`
	Game_Name string `json:"game_name"`
	Winner    string `json:"winner"`
	Queue     string `json:"queue"`
}

type Mongo struct { // struct para credenciales de mongo
	Database   string `json:"database"`
	Collection string `json:"collection"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Url        string `json:"url"`
}

type Redis struct { // struct para credenciales de redis y tidis
	Url  string `json:"url"`
	Type string `json:"type"`
}

var ( // se inicializan varialbes globales con credenciales para las db
	_redis = Redis{
		fmt.Sprintf("%v:%v", Getenv("REDIS_HOST"), Getenv("REDIS_PORT")),
		"Redis",
	}

	_tidis = Redis{
		fmt.Sprintf("%v:%v", Getenv("TIDIS_HOST"), Getenv("TIDIS_PORT")),
		"Tidis",
	}

	_mongo = Mongo{
		Getenv("MONGO_DB"),
		Getenv("MONGO_COL"),
		Getenv("MONGO_USER"),
		Getenv("MONGO_PASS"),
		fmt.Sprintf(`mongodb://%v:%v/?authSource=admin&readPreference=primary&directConnection=true&ssl=false`, os.Getenv("MONGO_HOST"), os.Getenv("MONGO_PORT")),
	}
)

// funcion para verificar variables de entorno
func Getenv(key string) string {
	value, defined := os.LookupEnv(key)
	if !defined {
		errnv := godotenv.Load()
		if errnv != nil {
			return ""
		} else {
			return Getenv(key)
		}
	}
	return value
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

// Funcion para obtener la data de la cola

func getMsg(msg []byte) {
	var dataRabbit getRabbitData
	// convertimos la data a json
	errJason := bson.UnmarshalExtJSON(msg, true, &dataRabbit)
	if errJason != nil {
		fmt.Println(errJason)
	}
	// imprimimos los valores obtenidos
	fmt.Println("Game id: ", dataRabbit.Game_Id)
	fmt.Println("Players: ", dataRabbit.Players)
	fmt.Println("Game_name: ", dataRabbit.Game_Name)
	fmt.Println("Winner: ", dataRabbit.Winner)
	fmt.Println("Queue: ", dataRabbit.Queue)

	mongoSaver(context.Background(), dataRabbit, _mongo)
	rtDisSaver(context.Background(), dataRabbit, _redis)
	rtDisSaver(context.Background(), dataRabbit, _tidis)
}

// Funcion para insertar en mongo
func mongoSaver(ctx context.Context, match getRabbitData, _mongo Mongo) {

	credential := options.Credential{
		Username: _mongo.Username,
		Password: _mongo.Password,
	}

	ctxMongo, cancel := context.WithTimeout(ctx, time.Second*10)
	clientOptions := options.Client().ApplyURI(_mongo.Url).SetAuth(credential)

	c, clientError := mongo.NewClient(clientOptions)
	if clientError != nil {
		fmt.Println("Mongo: " + clientError.Error())
	}

	connectError := c.Connect(ctxMongo)
	if connectError != nil {
		fmt.Println("Mongo: " + connectError.Error())
	}

	pingError := c.Ping(ctxMongo, nil)
	if pingError != nil {
		fmt.Println("Mongo: " + pingError.Error())
	}

	ctxInsert := context.Background()
	todoCollection := c.Database(_mongo.Database).Collection(_mongo.Collection)
	_, insertError := todoCollection.InsertOne(ctxInsert, match)

	if insertError != nil {
		fmt.Println("Mongo: " + insertError.Error())
	} else {
		fmt.Println("Mongo: Saved")
	}

	disconnectError := c.Disconnect(ctxInsert)
	if disconnectError != nil {
		return
	}
	cancel()
}

// Funcion para insertar en Redis y Tidis
func rtDisSaver(ctx context.Context, match getRabbitData, _rtDis Redis) {
	client := redis.NewClient(&redis.Options{
		Addr: _rtDis.Url,
		DB:   0,
	})
	data, encoderError := encoder.Marshal(match)
	if encoderError != nil {
		fmt.Println(_rtDis.Type+": ", encoderError.Error())
	}
	client.RPush(ctx, "logskafka", data)
	fmt.Println(_rtDis.Type + ": Saved")
}

func main() {
	conn, err := amqp.Dial(Getenv("Prot_Rabbit") + "://" + Getenv("Us_Rabbit") + ":" + Getenv("Ps_Rabbit") + "@" + Getenv("Host_Rabbit") + ":" + Getenv("Port_Rabbit") + "/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"fase2", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	var forever chan struct{}

	go func() {
		for d := range msgs {
			getMsg(d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever

}
