package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	pb "github.com/AllVides/so1_proyecto/fase2/gRPC-client-api/proto"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
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
	w.Write([]byte("Proyecto Fase 3 - Grupo 2\n"))
}

// funcion para ejecutar gRCP
func RunPlay(w http.ResponseWriter, r *http.Request) {
	var data dataJson

	/************ Leemos datos enviados desde nuestro CLI ***************/

	queryGet := r.URL.Query() // variable que nos ayudara a obtener los parametros de la URL

	idGame := queryGet.Get("game_id")      // obtenemos el id del juego
	intId, errConv := strconv.Atoi(idGame) // convertimos id en entero de 32 bits
	if errConv != nil {
		log.Fatal(">> gRPC-Client-id del juego no es número")
	}
	data.Game_id = int32(intId)

	playersGame := queryGet.Get("players")           // obtenemos el total de jugadores
	intPlayers, errConv := strconv.Atoi(playersGame) // convertimos el total de jugadores en entero de 32 bits
	if errConv != nil {
		log.Fatal(">> gRPC-Client-total de jugadores no es número")
	}
	data.Players = int32(intPlayers)

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
	router.HandleFunc("/on", IndexHandler)
	router.HandleFunc("/", RunPlay).Methods("GET")
	log.Println("Listening at port 3000")
	log.Fatal(http.ListenAndServe(":3000", handlers.CORS(headers, methods, origins)(router)))
}
