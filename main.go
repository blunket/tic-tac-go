package main

import (
	"log"
	"net"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "54321"
	CONN_TYPE = "tcp"
)

func main() {

	ln, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()
	log.Printf("Listening on %s:%s", CONN_HOST, CONN_PORT)
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	buf := make([]byte, 1024)

	_, err := conn.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	conn.Write([]byte("Received message."))
	conn.Close()
}
