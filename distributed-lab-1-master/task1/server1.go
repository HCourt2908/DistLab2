package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {

	//initialises the server
	ln, _ := net.Listen("tcp", ":8030")

	for {
		//accepts connection from new client
		conn, _ := ln.Accept()

		//reads and outputs message from the client
		reader := bufio.NewReader(conn)
		msg, _ := reader.ReadString('\n')
		fmt.Println("Client says: " + msg)

		//sends the response message to the client
		fmt.Fprintln(conn, "Server received message: "+msg)
	}
}
