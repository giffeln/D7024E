package main

import (
	"bytes"
	"encoding/gob"
	"flag"
	"fmt"
	"net"
)

var ipFlag = flag.String("i", "localhost", "ip of node you want to message")
var portFlag = flag.String("p", "3000", "port of the node you want to message")
var idFlag = flag.String("c", "PING", "command to send")
var dataFlag = flag.String("d", "", "data to send")

type Message struct {
	ID   string
	Data string
}

func client(ip string, port string, msg Message) {
	ipPort := ip + ":" + port
	//conn, _ := net.Dial("tcp", ipPort)
	conn, _ := net.Dial("tcp", ipPort)

	// Uncomment to test timeout
	// time.Sleep(5 * time.Second)
	// return

	send(conn, msg)
	recv(conn)
}

func send(conn net.Conn, msg Message) {
	// lets create the message we want to send accross
	//msg := Message{ID: "Yo", Data: "Hello"}
	bin_buf := new(bytes.Buffer)

	// create a encoder object
	gobobj := gob.NewEncoder(bin_buf)
	// encode buffer and marshal it into a gob object
	gobobj.Encode(msg)

	conn.Write(bin_buf.Bytes())
}

func recv(conn net.Conn) {
	// create a temp buffer
	tmp := make([]byte, 500)
	conn.Read(tmp)

	// convert bytes into Buffer (which implements io.Reader/io.Writer)
	tmpbuff := bytes.NewBuffer(tmp)
	tmpstruct := new(Message)

	// creates a decoder object
	gobobjdec := gob.NewDecoder(tmpbuff)
	// decodes buffer and unmarshals it into a Message struct
	gobobjdec.Decode(tmpstruct)

	fmt.Println(tmpstruct)
}

func main() {
	flag.Parse()
	var msg = Message{ID: *idFlag, Data: *dataFlag}
	client(*ipFlag, *portFlag, msg)
}
