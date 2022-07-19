package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	pb "github.com/AllVides/so1_proyecto/fase2/gRPC-client-api/proto"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// structs para obtener los datos del juego
type dataJson struct {
	Game_id int32 `json:"game_id"`
	Players int32 `json:"players"`
}

// Funcion que establece cabeceras de respuesta
func PushHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
}

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

// funcion cuando se hace peticion a la raiz
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Proyecto Fase 2 - Grupo 2\n"))
}

// funcion para ejecutar gRCP
func RunPlay(w http.ResponseWriter, r *http.Request) {
	/************ Leemos datos enviados desde locust ***************/
	//PARCEO EL JSON QUE ME ENVIARON
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	str := buf.String()

	var data dataJson
	err := bson.UnmarshalExtJSON([]byte(str), true, &data) // convertimos la data recibida a json
	if err != nil {
		fmt.Println(err)
	}

	/********************** gRPC **********************************/
	portServer := Getenv("Host_gRPCServ") + ":" + Getenv("Port_gRPCServ")
	// Set up a connection to the server.
	conn, err := grpc.Dial(portServer, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf(">> gRPC-Client-no se ha podido conectar: %v", err)
	}
	defer conn.Close()
	c := pb.NewRunPlayClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	reply, err := c.SenConfigurationPlay(ctx, &pb.PlayRequest{
		GameId: data.Game_id,
		Player: data.Players,
	})
	if err != nil {
		log.Fatalf(">> gRPC-Client-no se ha podido jugar: %v", err)
	}
	msj := reply.GetMessage()
	log.Printf(">> gRPC-Client-jugando: %s", msj)

	// enviamos struct de data leida
	PushHeaders(w)
	json.NewEncoder(w).Encode(msj)
}

// funcion principal
func main() {
	router := mux.NewRouter().StrictSlash(false)
	headers := handlers.AllowedHeaders([]string{"X-Request-Width", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST"})
	origins := handlers.AllowedOrigins([]string{"*"})
	router.HandleFunc("/", IndexHandler)
	router.HandleFunc("/play", RunPlay).Methods("POST")
	log.Println("Listening at port 8080")
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(headers, methods, origins)(router)))
}
