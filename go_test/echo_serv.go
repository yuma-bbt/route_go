package main

import (
	"io"
	"log"
	"net"
)

func main() {
	url := "localhost"
	port := "8888"

	listener, err := net.Listen("tcp", url+":"+port)
	if err != nil {
		panic(err)
	}
	log.Println("Server start at ", url+":"+port)
	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		go func() {
			defer conn.Close()
			echo(conn)
		}()
	}
}

func echo(conn net.Conn) {
	messageBuf := make([]byte, 256)
	for {
		messageLength, err := conn.Read(messageBuf)
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		if messageLength == 0 {
			break
		}
		log.Println("Request", string(messageBuf))
		_, err = conn.Write(messageBuf)
		if err != nil {
			panic(err)
		}
		log.Println("Respoince  ", string(messageBuf))
	}
}
