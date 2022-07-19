package TestGCP

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"time"
)

// Database connection
var (
	usr      = "sistemasOperativos1"
	pwd      = "1234"
	host     = "34.67.195.168"
	port     = 27017
	database = "projectDatabase"
)

func GetCollection(collection string) *mongo.Collection {
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%d", usr, pwd, host, port)
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))

	if err != nil {
		panic(err.Error())
	}

	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)

	if err != nil {
		panic(err.Error())
	}

	return client.Database(database).Collection(collection)
}

// RamLog Type
type RamLog struct {
	NombreVM string    `json:"nombreVM"`
	Endpoint string    `json:"endpoint"`
	Data     []RamData `json:"data"`
	Date     string    `json:"date"`
}

type RamData struct {
	Total        string `json:"total"`
	MemoriaEnUso string `json:"memoriaEnUso"`
	Porcentaje   string `json:"porcentaje"`
	MemoriaLibre string `json:"memoriaLibre"`
}

type ProcLog struct {
	Procs    []Proceso `json:"procs"`
	NombreVM string    `json:"nombreVM"`
	Date     string    `json:"date"`
	Endpoint string    `json:"endpoint"`
}

type Proceso struct {
	Pid    string    `json:"pid"`
	Nombre string    `json:"nombre"`
	Estado string    `json:"estado"`
	Hijo   []Proceso `json:"hijo"`
}

// Repository
var (
	collection = GetCollection("projectDatabase")
	ctx        = context.Background()
)

func CreateRam(ramLog RamLog) error {
	var err error

	_, err = collection.InsertOne(ctx, ramLog)

	if err != nil {
		return err
	}

	return nil
}

func CreateProc(procLog ProcLog) error {
	var err error

	_, err = collection.InsertOne(ctx, procLog)

	if err != nil {
		return err
	}

	return nil
}

func main() {
	router := mux.NewRouter()
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})
	router.HandleFunc("/createRamLog", CreateRamLog).Methods("POST")
	router.HandleFunc("/createProcLog", CreateProcLog).Methods("POST")
	err := http.ListenAndServe(":12345", handlers.CORS(headers, methods, origins)(router))
	if err != nil {
		return
	}
}

func CreateRamLog(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	var ramLog RamLog
	err := json.NewDecoder(request.Body).Decode(&ramLog)
	if err != nil {
		json.NewEncoder(response).Encode("Error while decoding body, check types")
		return
	}
	err = CreateRam(ramLog)
	if err != nil {
		err = json.NewEncoder(response).Encode("Error saving log")
		return
	} else {
		err = json.NewEncoder(response).Encode(ramLog)
	}
}

func CreateProcLog(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	var procLog ProcLog
	err := json.NewDecoder(request.Body).Decode(&procLog)
	if err != nil {
		json.NewEncoder(response).Encode("Error while decoding body, check types")
		return
	}
	err = CreateProc(procLog)
	if err != nil {
		err = json.NewEncoder(response).Encode("Error saving log")
		return
	} else {
		err = json.NewEncoder(response).Encode(procLog)
	}
}
