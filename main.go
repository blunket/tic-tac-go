package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	// "github.com/blunket/tic-tac-go/game"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "54321"
)

func main() {

	r := httprouter.New()

	r.GET("/new", NewGame)

	r.GET("/grid", Grid)

	log.Printf("Listening on %s:%s\n", CONN_HOST, CONN_PORT)
	log.Fatalln(http.ListenAndServe(CONN_HOST+":"+CONN_PORT, r))

}

func Grid(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Hello, world!")
	log.Println("/grid")
}

func NewGame(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "New game!")
	log.Println("/new")
}
