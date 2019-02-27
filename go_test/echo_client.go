package main

import (
	"log"
	"net"
)

func main() {
	url := "localhost"
	port := "8888"
	conn, err := net.Dial("tcp", url+":"+port)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	log.Println("Connected to server ")
	text := "hogehoge"
	_, err = conn.Write([]byte(text))

	messageBuf := make([]byte, 256)
	n, err := conn.Read(messageBuf)
	if err != nil {
		panic(err)
	}
	log.Println("Server : ", string(messageBuf[:n]))
}
