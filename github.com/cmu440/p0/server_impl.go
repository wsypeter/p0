// Implementation of a KeyValueServer. Students should write their code in this file.

package p0
import (
	"fmt"
	"net"
	"io"
	// "strconv"
)
type client struct {
	name string
	connection net.Conn
	messageQueue chan []byte
}
type db struct {
	isGet bool
	key []byte
	value []byte
}
type keyValueServer struct {
	connection chan net.Conn
	clientN chan int
	clientList [] *client
	listener net.Listener
	deadClient chan *client
	newMessage chan []byte

		// TODO: implement this!
}

// New creates and returns (but does not start) a new KeyValueServer.
func New() KeyValueServer {
    // TODO: implement this!
    return &keyValueServer{
		make (chan net.Conn),
		make (chan int),
		make ([] *client, 0),
		nil,
		make (chan *client),
		make (chan []byte)}
}

func (kvs *keyValueServer) Start(port int) error {
	ln, err := net.Listen("tcp", ":9999")
	if err != nil {
		// handle error
	}
	kvs.listener = ln;
	init_db()
	go acceptRoutine(kvs);
	go runServer(kvs)
	return nil	
}

func acceptRoutine(kvs *keyValueServer){
	for {
		conn, err :=kvs.listener.Accept();
		fmt.Println(conn)
		if err == nil {
			fmt.Println("writing to connect channel")
			kvs.connection <- conn;
			fmt.Println(kvs.connection)
		}
	}
}

func runServer(kvs *keyValueServer) {
	fmt.Println("in runserver");
	for {
		select {
		case newMessage := <- kvs.newMessage:
			for _, v := range kvs.clientList {
				v.messageQueue <- newMessage
			}
		case newConn := <-kvs.connection:
			fmt.Println("new connection")
			c := &client{
				"testclient",
				newConn,
				make(chan []byte) }
			kvs.clientList = append(kvs.clientList, c)
			go readRoutine(kvs, c)
			go writeRoutine( c)

		case deadClient := <-kvs.deadClient:
			for i, c := range kvs.clientList {
				if c == deadClient {
					fmt.Println("removing dead")
					kvs.clientList =
						append(kvs.clientList[:i], kvs.clientList[i+1:]...)
					break
				}
			}
		case query := kvs.queryChan
			if query.get {
				v := go get (request.key)
				kvs.response <- v

			} else {
				go put (request.key, request.value)
			}
		}		

	}
}

func readRoutine(kvs *keyValueServer, c *client) {
	
  // Make a buffer to hold incoming data.
  clientReader := bufio.NewReader(c.connection)
  for {
	message, err := clientReader.ReadBytes('\n')
	if err == io.EOF {
		kvs.deadClient <- c; 
	} else if err != nil {
		fmt.Println(err)
		return
	} else {// processsing
		buffer := bytes.Split(message, []byte(","))
		if (string(buffer[0]) == "put") {
			key := string (buffer[1])
			value := string (buffer[2])
			//do db query

		} else{ //read
			key := string (buffer[1])
			kvs.queryChan <- &db{
				isGet: true,
				key:key,
			}
			response := kvs.response
			kvs.newMessage <- response

		}
		fmt.Println(string(buf));  // Send a response back to person contacting us.
		kvs.newMessage <- buf;
	}
  }
  // Read the incoming connection into the buffer.
}
func writeRoutine (c *client){
	for {
		<- c.messageQueue
		c.connection.Write([]byte("roger\n"))
	}
}



func (kvs *keyValueServer) Close() {
    // TODO: implement this!
}

func (kvs *keyValueServer) Count() int {
    // TODO: implement this!
    return -1
}

// TODO: add additional methods/functions below!