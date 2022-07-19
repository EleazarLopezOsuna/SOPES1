package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	// para conectar con cosmosDB
)

// structs para data ram
type ramJson struct { // struct para obtener data del proc ram.ko
	MemoriaTotal      string `json:"total"`
	MemoriaLibre      string `json:"memoriaLibre"`
	MemoriaUso        string `json:"memoriaEnUso"`
	MemoriaUsoPercent string `json:"porcentaje"`
}

type dataRamJson struct {
	Data           []ramJson `json:"data"`
	VirtualMachine string    `json:"nombreVM"`
	DateQuery      string    `json:"date"`
	EndpointName   string    `json:"endpoint"`
}

type returnJson struct {
	RamJson dataRamJson `json:"ramJson"`
	IsError bool        `json:"isError"`
}

type responseJsonRam struct {
	Data    []ramJson `json:"data"`
	MV      string    `json:"mv"`
	IsError bool      `json:"isError"`
}

// structs para data procesos
type procesosJson struct {
	Pid    string         `json:"pid"`
	Nombre string         `json:"nombre"`
	Estado string         `json:"estado"`
	Hijo   []procesosJson `json:"hijo"`
}

type dataProcesosJson struct { // struct para obtener data del proc procesos.ko
	Procs          []procesosJson `json:"procs"`
	VirtualMachine string         `json:"nombreVM"`
	DateQuery      string         `json:"date"`
	EndpointName   string         `json:"endpoint"`
}

type returnProcesosJson struct {
	Procesos dataProcesosJson `json:"procesos"`
	IsError  bool             `json:"isError"`
}

type responseJsonProc struct {
	Data    []procesosJson `json:"data"`
	MV      string         `json:"mv"`
	IsError bool           `json:"isError"`
}

// return para endpoint /
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Proyecto Fase 1 - Grupo 2\n"))
}

/******************************************* Para obtener datos de RAM ***********************************************/
// return para endpoint /ram1
func GetHandlerRam1(w http.ResponseWriter, r *http.Request) {
	var dataResponse responseJsonRam
	vm := Getenv("VM")
	if vm == "" {
		vm = "1"
	}
	data := GetRam(vm)
	data.RamJson.EndpointName = "/ram1"

	// enviamos struct de data leida
	PushHeaders(w)
	if !data.IsError {
		dataResponse.Data = data.RamJson.Data
		dataResponse.MV = data.RamJson.VirtualMachine
		dataResponse.IsError = SendRam(data.RamJson)
	} else {
		dataResponse.IsError = true
	}
	json.NewEncoder(w).Encode(dataResponse)
}

// return para endpoint /ram2
func GetHandlerRam2(w http.ResponseWriter, r *http.Request) {
	var dataResponse responseJsonRam
	vm := Getenv("VM")
	if vm == "" {
		vm = "2"
	}
	data := GetRam(vm)
	data.RamJson.EndpointName = "/ram2"

	// enviamos struct de data leida
	PushHeaders(w)
	if !data.IsError {
		dataResponse.Data = data.RamJson.Data
		dataResponse.MV = data.RamJson.VirtualMachine
		dataResponse.IsError = SendRam(data.RamJson)
	} else {
		dataResponse.IsError = true
	}
	json.NewEncoder(w).Encode(dataResponse)
}

// funcion para obtener la ram de la maquina
func GetRam(vm string) returnJson {
	var ramReturn returnJson
	var dataRam ramJson
	var data dataRamJson
	isError := false

	// leemos el archivo creado en nuestros procs
	dataFileRam, err := ioutil.ReadFile(Getenv("PATH_RAM"))
	if err != nil {
		fmt.Print(err)
		isError = true
	}

	// convertimos la data a json
	errJason := bson.UnmarshalExtJSON(dataFileRam, true, &dataRam)
	if errJason != nil {
		fmt.Println(errJason)
		isError = true
	}

	// quitamos espacios demás
	dataRam.MemoriaTotal = strings.TrimSpace(dataRam.MemoriaTotal)
	dataRam.MemoriaUso = strings.TrimSpace(dataRam.MemoriaUso)
	dataRam.MemoriaLibre = strings.TrimSpace(dataRam.MemoriaLibre)

	// convertimos a enteros la memoria total y la memoria en uso
	intMemoriaTotal, errMemoTotal := strconv.Atoi(dataRam.MemoriaTotal)
	intMemoriaUso, errMemoUso := strconv.Atoi(dataRam.MemoriaUso)
	if errMemoTotal != nil {
		fmt.Println(errMemoTotal)
		isError = true
	} else if errMemoUso != nil {
		fmt.Println(errMemoUso)
		isError = true
	}

	// obtenemos el porcentaje de memoria en uso y lo convertimos a string
	dataRam.MemoriaUsoPercent = strconv.Itoa(int(float64(intMemoriaUso) / float64(intMemoriaTotal) * 100))

	// Indicamos de que maquina virtual son los datos
	data.VirtualMachine = vm

	// imprimimos los valores obtenidos
	fmt.Println("Memoria total maquina virtual "+vm+": ", dataRam.MemoriaTotal)
	fmt.Println("Memoria en Uso maquina virtual "+vm+": ", dataRam.MemoriaUso)
	fmt.Println("Memoria libre maquina virtual "+vm+": ", dataRam.MemoriaLibre)
	fmt.Println("Porcentaje de memoria en uso maquina virtual "+vm+": ", dataRam.MemoriaUsoPercent)
	fmt.Println("Maquina virtual " + vm)

	data.DateQuery = GetTime()
	data.Data = append(data.Data, dataRam)
	ramReturn.RamJson = data
	ramReturn.IsError = isError

	return ramReturn
}

// funcion para enviar proc/ram_G2 al log
func SendRam(data dataRamJson) bool {
	var isError = false
	clienteHttp := &http.Client{}

	jsonData, err := json.Marshal(data)
	if err != nil {
		isError = true
		fmt.Println("Error codificando data como JSON: ", err)
	}

	peticion, err := http.NewRequest("POST", "https://"+Getenv("GCF_RAM"), bytes.NewBuffer(jsonData)) // generamos la peticion
	if err != nil {
		isError = true
		fmt.Println("Error creando petición: ", err)
	}

	peticion.Header.Add("Content-Type", "application/json") // añadimos el tipo de contenido que se enviara
	respuesta, err := clienteHttp.Do(peticion)              // realizamos la peticion
	if err != nil {
		isError = true
		fmt.Println("Error al realizar la petición: ", err)
	}

	defer respuesta.Body.Close() // cerramos la peticion

	return isError
}

/******************************************* Para obtener datos de los procesos ***********************************************/
// return para endpoint /procesos1
func GetHandlerProcess1(w http.ResponseWriter, r *http.Request) {
	var dataResponse responseJsonProc
	vm := Getenv("VM")
	if vm == "" {
		vm = "1"
	}

	data := GetProcess(vm)
	data.Procesos.EndpointName = "/procesos1"

	// enviamos struct de data leida
	PushHeaders(w)
	if !data.IsError {
		dataResponse.Data = data.Procesos.Procs
		dataResponse.MV = data.Procesos.VirtualMachine
		dataResponse.IsError = SendProcs(data.Procesos)
	} else {
		dataResponse.IsError = true
	}
	json.NewEncoder(w).Encode(dataResponse)
}

// return para endpoint /procesos2
func GetHandlerProcess2(w http.ResponseWriter, r *http.Request) {
	var dataResponse responseJsonProc
	vm := Getenv("VM")
	if vm == "" {
		vm = "2"
	}

	data := GetProcess(vm)
	data.Procesos.EndpointName = "/procesos2"

	// enviamos struct de data leida
	PushHeaders(w)
	if !data.IsError {
		dataResponse.Data = data.Procesos.Procs
		dataResponse.MV = data.Procesos.VirtualMachine
		dataResponse.IsError = SendProcs(data.Procesos)
	} else {
		dataResponse.IsError = true
	}
	json.NewEncoder(w).Encode(dataResponse)
}

// funcion para obtener los procesos de la maquina
func GetProcess(vm string) returnProcesosJson {
	var dataProces []procesosJson
	var data dataProcesosJson
	var processReturn returnProcesosJson
	isError := false

	// leemos el archivo creado en nuestros procs
	dataFileProcess, err := ioutil.ReadFile(Getenv("PATH_PROC"))
	if err != nil {
		fmt.Print(err)
		isError = true
	}

	// convertimos la data a json
	errJason := bson.UnmarshalExtJSON(dataFileProcess, true, &dataProces)
	if errJason != nil {
		fmt.Println(errJason)
		isError = true
	}

	data.Procs = dataProces

	// Indicamos de que maquina virtual son los datos
	data.VirtualMachine = vm
	data.DateQuery = GetTime()

	processReturn.Procesos = data
	processReturn.IsError = isError

	return processReturn
}

// funcion para enviar proc/procesos_G2 al log
func SendProcs(data dataProcesosJson) bool {
	var isError = false
	clienteHttp := &http.Client{}

	jsonData, err := json.Marshal(data)
	if err != nil {
		isError = true
		fmt.Println("Error codificando data como JSON: ", err)
	}

	peticion, err := http.NewRequest("POST", "https://"+Getenv("GCF_PROCS"), bytes.NewBuffer(jsonData)) // generamos la peticion
	if err != nil {
		isError = true
		fmt.Println("Error creando petición: ", err)
	}

	peticion.Header.Add("Content-Type", "application/json") // añadimos el tipo de contenido que se enviara
	respuesta, err := clienteHttp.Do(peticion)              // realizamos la peticion
	if err != nil {
		isError = true
		fmt.Println("Error al realizar la petición: ", err)
	}

	defer respuesta.Body.Close() // cerramos la peticion

	return isError
}

/********************************************** Otras funciones *****************************************************/
// funcion para obtener la hora y fecha actual
func GetTime() string {
	t := time.Now()
	fecha := fmt.Sprintf("%02d-%02d-%d %02d:%02d:%02d",
		t.Day(), t.Month(), t.Year(),
		t.Hour(), t.Minute(), t.Second())

	return fecha
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

// Funcion que establece cabeceras de respuesta
func PushHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
}

// funcion principal
func main() {
	router := mux.NewRouter().StrictSlash(false)
	headers := handlers.AllowedHeaders([]string{"X-Request-Width", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST"})
	origins := handlers.AllowedOrigins([]string{"*"})
	router.HandleFunc("/", IndexHandler)
	router.HandleFunc("/ram1", GetHandlerRam1).Methods("GET")
	router.HandleFunc("/ram2", GetHandlerRam2).Methods("GET")
	router.HandleFunc("/procesos1", GetHandlerProcess1).Methods("GET")
	router.HandleFunc("/procesos2", GetHandlerProcess2).Methods("GET")
	log.Println("Listening at port 8080")
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(headers, methods, origins)(router)))
}
