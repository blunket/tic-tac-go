package main

import (
	"bufio"
	"log"
	"net"
	// "github.com/blunket/tic-tac-go/game"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "54321"
	CONN_TYPE = "tcp"
)

func main() {

	ln, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	defer ln.Close()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Listening on %s:%s", CONN_HOST, CONN_PORT)

	for {
		conn, err := ln.Accept()
		if err != nil {
			break
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {

	defer conn.Close()
	for {
		msg, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			break
		}
		conn.Write([]byte("Received message: " + msg + "\n"))
		log.Println(msg)
	}

}
