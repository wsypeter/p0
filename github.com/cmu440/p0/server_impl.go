// Implementation of a KeyValueServer. Students should write their code in this file.

package p0
import (
	"fmt"
	"net"
	// "strconv"
)

type keyValueServer struct {
		// TODO: implement this!
}

// New creates and returns (but does not start) a new KeyValueServer.
func New() KeyValueServer {
    // TODO: implement this!
    return &keyValueServer{};
}

func (kvs *keyValueServer) Start(port int) error {
	ln, err := net.Listen("tcp", ":9999")
	if err != nil {
		// handle error
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err);
			// handle error
		}
		go handleRequest(conn);
	}
}
func handleRequest(conn net.Conn) {
  // Make a buffer to hold incoming data.
  buf := make([]byte, 1024)
  // Read the incoming connection into the buffer.
  _, err := conn.Read(buf)
  if err != nil {
    fmt.Println("Error reading:", err.Error())
	}
	fmt.Println(string(buf));  // Send a response back to person contacting us.
  conn.Write([]byte("Message received."))
  // Close the connection when you're done with it.
  conn.Close()
}

func (kvs *keyValueServer) Close() {
    // TODO: implement this!
}

func (kvs *keyValueServer) Count() int {
    // TODO: implement this!
    return -1
}

// TODO: add additional methods/functions below!