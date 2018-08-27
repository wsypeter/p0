package main

import (
	"net"
	"bufio"
	"strconv"
	"fmt"
)

const (
    defaultHost = "localhost"
    defaultPort = 9999
)

// To test your server implementation, you might find it helpful to implement a
// simple 'client runner' program. The program could be very simple, as long as
// it is able to connect with and send messages to your server and is able to
// read and print out the server's response to standard output. Whether or
// not you add any code to this file will not affect your grade.
func main() {
    conn, _ := net.Dial("tcp", ":" + strconv.Itoa(defaultPort))

	for i := 0; i < 100; i++ {
		message := strconv.Itoa(i) + "\n"
		conn.Write([]byte(message))
	}
	for {
		// will listen for message to process ending in newline (\n)
		message, _ := bufio.NewReader(conn).ReadString('\n')
		// output message received
		fmt.Print("Message Received:", string(message))
	  }

	// select{}
}
