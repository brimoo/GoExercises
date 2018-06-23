/*
	Written while following along with the talk "Go: code that grows with grace". Check out the talk
	here: https://vimeo.com/53221560
*/

package main

import (
	"net"
	"log"
	"fmt"
)

const listenerAddr = "localhost:4000"

func main() {
	listener, err := net.Listen("tcp", listenerAddr)
	if err != nil {
		log.Fatal(err)
	}
	for {
		connection, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprintln(connection, "Hello!")
		connection.Close()
	}
}