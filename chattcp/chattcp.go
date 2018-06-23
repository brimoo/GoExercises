/*
	Written while following along with the talk "Go: code that grows with grace". Check out the talk
	here: https://vimeo.com/53221560
*/

package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

const listenerAddr = "localhost:4000"

var partnerChannel = make(chan io.ReadWriteCloser)

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
		go findChatPartner(connection)
	}
}

func findChatPartner(connection io.ReadWriteCloser) {
	fmt.Fprintln(connection, "Waiting for partner...")
	select {
	case partner := <-partnerChannel:
		chat(connection, partner)
	case partnerChannel <- connection:
		// Partner will handle the connection
	}
}

func chat(a, b io.ReadWriteCloser) {
	fmt.Fprintln(a, "Found partner. Say something!")
	fmt.Fprintln(b, "Found partner. Say something!")
	errorChannel := make(chan error, 1)
	go sendMessage(a, b, errorChannel)
	go sendMessage(b, a, errorChannel)
	if err := <-errorChannel; err != nil {
		log.Println(err)
	}
	a.Close()
	b.Close()
}

func sendMessage(writer io.Writer, reader io.Reader, errorChannel chan<- error) {
	_, err := io.Copy(writer, reader)
	errorChannel <- err
}
