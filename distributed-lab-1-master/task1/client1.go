package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {

	for {
		//establish connection with server
		conn, _ := net.Dial("tcp", "127.0.0.1:8030")

		//get user input
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("Enter text")
		scanner.Scan()
		text := scanner.Text()

		//send message to the server
		fmt.Fprintln(conn, text)

		//receives and outputs the response from the server
		reader := bufio.NewReader(conn)
		msg, _ := reader.ReadString('\n')
		fmt.Println(msg)

	}
}
