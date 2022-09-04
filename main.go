//a simple tcp server that echo a given response back to the client
package main

import (
	"bufio"
	"fmt"
	"net"

)

func main() {
	// listen on all interfaces
	ln, _ := net.Listen("tcp", ":8081")

	// accept connection on port
	conn, _ := ln.Accept()

	// run loop forever (or until ctrl-c)
	for {
		// will listen for message to process ending in newline (\n)
		message, _ := bufio.NewReader(conn).ReadString('\n')
		// output message received
		fmt.Print("Message Received:", string(message))
		// sample process for string received
		newmessage := message
		// send new string back to client
		conn.Write([]byte(newmessage + ""))
	}
}

