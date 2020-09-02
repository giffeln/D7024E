package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

// Create your custom data struct
type Message struct {
	ID   string
	Data string
}

func logerr(err error) bool {
	if err != nil {
		if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
			log.Println("read timeout:", err)
		} else if err == io.EOF {
		} else {
			log.Println("read error:", err)
		}
		return true
	}
	return false
}

func server() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		// handle error
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
		}
		go handleConnection(conn)
	}
}

func read(conn net.Conn) {
	// create a temp buffer
	tmp := make([]byte, 500)

	// loop through the connection to read incoming connections. If you're doing by
	// directional, you might want to make this into a seperate go routine
	for {
		_, err := conn.Read(tmp)
		if logerr(err) {
			break
		}

		// convert bytes into Buffer (which implements io.Reader/io.Writer)
		tmpbuff := bytes.NewBuffer(tmp)
		tmpstruct := new(Message)

		// creates a decoder object
		gobobj := gob.NewDecoder(tmpbuff)
		// decodes buffer and unmarshals it into a Message struct
		gobobj.Decode(tmpstruct)

		// lets print out!
		fmt.Println(tmpstruct)
		return
	}
}

func resp(conn net.Conn) {
	msg := Message{ID: "Yo", Data: "Hello back"}
	bin_buf := new(bytes.Buffer)

	// create a encoder object
	gobobje := gob.NewEncoder(bin_buf)
	// encode buffer and marshal it into a gob object
	gobobje.Encode(msg)

	conn.Write(bin_buf.Bytes())
	conn.Close()
}

func handleConnection(conn net.Conn) {
	timeoutDuration := 2 * time.Second
	fmt.Println("Launching server...")
	conn.SetReadDeadline(time.Now().Add(timeoutDuration))

	remoteAddr := conn.RemoteAddr().String()
	fmt.Println("Client connected from " + remoteAddr)

	read(conn)
	resp(conn)
}

func client() {
	conn, _ := net.Dial("tcp", ":8080")

	// Uncomment to test timeout
	// time.Sleep(5 * time.Second)
	// return

	send(conn)
	recv(conn)
}

func send(conn net.Conn) {
	// lets create the message we want to send accross
	msg := Message{ID: "Yo", Data: "Hello"}
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

	go server()

	for {
		time.Sleep(2 * time.Second)
		client()
	}
}
