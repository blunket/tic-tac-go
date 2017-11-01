package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	// "github.com/blunket/tic-tac-go/game"
)

const (
	connHost = "localhost"
	connPort = "54321"
)

func main() {

	r := httprouter.New()

	r.GET("/new", newGame)
	r.GET("/game/:gameId/grid", grid)

	log.Printf("Listening on %s:%s\n", connHost, connPort)
	log.Fatalln(http.ListenAndServe(connHost+":"+connPort, r))

}

func grid(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "Game: %s", ps.ByName("gameId"))
	log.Println("grid: ")
}

func newGame(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "New game!")
	log.Println("/new")
}
