package main

import (
	"context"
	encoder "encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/segmentio/kafka-go"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"time"
)

type Kafka struct {
	Broker string `json:"broker"`
	Topic  string `json:"topic"`
}

type Mongo struct {
	Database   string `json:"database"`
	Collection string `json:"collection"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Url        string `json:"url"`
}

type Redis struct {
	Url  string `json:"url"`
	Type string `json:"type"`
}

type Match struct {
	Game_Id   string `json:"game_id"`
	Players   string `json:"players"`
	Game_Name string `json:"game_name"`
	Winner    string `json:"winner"`
	Queue     string `json:"queue"`
}

func main() {

	/*envError := godotenv.Load(".env")
	if envError != nil {
		log.Fatal(envError.Error())
	}*/
	var (
		_kafka = Kafka{
			fmt.Sprintf("%v:%v", os.Getenv("KAFKA_HOST"), os.Getenv("KAFKA_PORT")),
			os.Getenv("KAFKA_TOPIC"),
		}

		_redis = Redis{
			fmt.Sprintf("%v:%v", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")),
			"Redis",
		}

		_tidis = Redis{
			fmt.Sprintf("%v:%v", os.Getenv("TIDIS_HOST"), os.Getenv("TIDIS_PORT")),
			"Tidis",
		}

		_mongo = Mongo{
			os.Getenv("MONGO_DB"),
			os.Getenv("MONGO_COL"),
			os.Getenv("MONGO_USER"),
			os.Getenv("MONGO_PASS"),
			fmt.Sprintf(`mongodb://%v:%v/?authSource=admin&readPreference=primary&directConnection=true&ssl=false`, os.Getenv("MONGO_HOST"), os.Getenv("MONGO_PORT")),
		}
	)

	ctxKafKa := context.Background()

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     []string{_kafka.Broker},
		Topic:       _kafka.Topic,
		GroupID:     "matches-group",
		StartOffset: kafka.LastOffset,
	})

	for {
		msg, readMessageError := r.ReadMessage(ctxKafKa)
		if readMessageError != nil {
			panic(readMessageError.Error())
		} else {
			var matchMap Match
			encoderError := encoder.Unmarshal(msg.Value, &matchMap)
			if encoderError != nil {
				fmt.Println("Error: ", encoderError.Error())
			}
			mongoSaver(context.Background(), matchMap, _mongo)
			rtDisSaver(context.Background(), matchMap, _redis)
			rtDisSaver(context.Background(), matchMap, _tidis)
		}
	}
}

func mongoSaver(ctx context.Context, match Match, _mongo Mongo) {

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

func rtDisSaver(ctx context.Context, match Match, _rtDis Redis) {
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
