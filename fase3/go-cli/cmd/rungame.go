/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/spf13/cobra"
)

var gamename, players, rungames, concurrence, timeout string

// obtenemos los id sin el nombre
func getId(listaGameName []string) []string {
	var ids []string

	for i := 0; i < len(listaGameName); i++ {
		idSplit := strings.Split(listaGameName[i], "|")
		ids = append(ids, strings.TrimSpace(idSplit[0]))
	}

	return ids
}

// obtenemos el tiempo y dimensional que se ejecutara la aplicación
func getTimeOutFlag(outime string) [2]string {
	numTime := ""
	dimensionalTime := "s"

	outime = strings.TrimSpace(outime)
	for i := 0; i < len(outime); i++ {

		_, errorConvert := strconv.Atoi(string(outime[i]))

		if errorConvert != nil {
			dimensionalTime = string(outime[i])
			retunrData := [2]string{numTime, dimensionalTime}
			return retunrData
		}

		numTime += string(outime[i])
	}
	retunrData := [2]string{numTime, dimensionalTime}
	return retunrData
}

// funcion que genera números random entre los players y total de ids
func generateRandom(min, max int, listIds []string) string {
	rand.Seed(time.Now().UnixNano())
	numRandom := rand.Intn(max-min) + min

	if listIds != nil {
		return listIds[numRandom]
	}
	return strconv.Itoa(numRandom)
}

// función que genera la petición a nuestro servicio
func generateGet(id, play string, waitGroup *sync.WaitGroup) {
	url := "http://34.72.66.110.nip.io?game_id=" + id + "&players=" + play
	fmt.Println(url)

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", body)

	time.Sleep(5 * time.Second)
	waitGroup.Done()
}

// funcion que generara la concurrencia
func generateConcurrence(concurrense, gamesrun, outime string, playrsInt int, listGameId []string) {
	randomId, randomPlayer := "", ""
	contadorPeticiones := 1
	gamesrunInt, _ := strconv.Atoi(gamesrun)
	concurrenseInt, _ := strconv.Atoi(concurrense)

	timeOutGame := getTimeOutFlag(outime) // Tiempo y Dimensional en que se ejecutara la aplicación
	timeFinishInt, _ := strconv.Atoi(timeOutGame[0])
	var timeStart time.Time

	if strings.ToLower(timeOutGame[1]) == "s" {
		timeStart = time.Now().Add(time.Second * time.Duration(timeFinishInt))
	} else if strings.ToLower(timeOutGame[1]) == "m" {
		timeStart = time.Now().Add(time.Minute * time.Duration(timeFinishInt))
	}

	waitGroup := &sync.WaitGroup{}

	for i := 0; i < concurrenseInt; i++ {
		timeFin := time.Now()
		if timeStart.After(timeFin) {
			randomPlayer = generateRandom(1, playrsInt, nil)          // Obtiene jugador random
			randomId = generateRandom(0, len(listGameId), listGameId) // Obtiene id random

			if contadorPeticiones <= gamesrunInt {
				waitGroup.Add(1)
				go generateGet(randomId, randomPlayer, waitGroup)
				contadorPeticiones++
			} else {
				fmt.Println("Se alcanzo el limite de endpoints utilizados")
				os.Exit(3)
			}
		} else {
			fmt.Println("Se alcanzo el limite de tiempo: ", timeFin)
			os.Exit(3)
		}

	}
	waitGroup.Wait()
}

// funcion encargada de realizar el proceso de la data recibida del CLI
func getData(namegame, playrs, gamesrun, concurrense, outime string) {

	fmt.Println(
		" Lista de juegos: ", namegame, "\n",
		"Cantidad de jugadores: ", players, "\n",
		"Cantidad de veces a jugar: ", rungames, "\n",
		"Cantidad de solicitudes al mismo tiempo: ", concurrence, "\n",
		"Tiempo de ejecución: ", timeout,
	)

	listaGameName := strings.Split(namegame, ";")
	playrsInt, _ := strconv.Atoi(strings.TrimSpace(playrs))

	listGameId := getId(listaGameName) // lista de IDS

	generateConcurrence(concurrense, gamesrun, outime, playrsInt, listGameId)
}

// rungameCmd represents the rungame command
var rungameCmd = &cobra.Command{
	Use:   "rungame",
	Short: "Generara varios juegos segun los parametros",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("rungame called \n")

		gamename, _ = cmd.Flags().GetString("gamename")
		players, _ = cmd.Flags().GetString("players")
		rungames, _ = cmd.Flags().GetString("rungames")
		concurrence, _ = cmd.Flags().GetString("concurrence")
		timeout, _ = cmd.Flags().GetString("timeout")

		getData(gamename, players, rungames, concurrence, timeout)

	},
}

func init() {
	rootCmd.AddCommand(rungameCmd)
	rungameCmd.PersistentFlags().StringVar(&gamename, "gamename", "", "lista de juegos a jugar")

	rungameCmd.PersistentFlags().StringVar(&players, "players", "", "Cantidad de Jugadores")
	rungameCmd.PersistentFlags().StringVar(&rungames, "rungames", "", "Cantidad de juegos a jugar")
	rungameCmd.PersistentFlags().StringVar(&concurrence, "concurrence", "", "Cantidad de gorutinas a ustilizar")
	rungameCmd.PersistentFlags().StringVar(&timeout, "timeout", "", "Cantidad de tiempo para abortar")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// rungameCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// rungameCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
