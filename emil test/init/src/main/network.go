package main

/*

source: https://www.linode.com/docs/development/go/developing-udp-and-tcp-clients-and-servers-in-go/

*/

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"math/rand"
	"net"
	"time"
)

const MESSAGE_SIZE int = 2048

type Contact struct {
	ID       byte
	Address  string
	Distance byte
}

type Message struct {
	CMD     string
	Data    string
	Contact Contact
}

/*
func messageToString(message Message) string {
	var newMessage string
	newMessage = message[ID] + "," + message["Data"]
	return newMessage
}

func stringToMessage(str string) Message {

}
*/

func client(ip string, port string, message Message) {
	s, err := net.ResolveUDPAddr("udp4", ip+":"+port)
	c, err := net.DialUDP("udp4", nil, s)
	if err != nil {
		fmt.Println(err)
		return
	}

	//fmt.Printf("The UDP server is %s\n", c.RemoteAddr().String())
	defer c.Close()

	/*reader := bufio.NewReader(os.Stdin)
	fmt.Print(">> ")
	text, _ := reader.ReadString('\n')*/

	bin_buf := new(bytes.Buffer)

	// create a encoder object
	gobobj := gob.NewEncoder(bin_buf)
	// encode buffer and marshal it into a gob object
	gobobj.Encode(message)

	_, err = c.Write(bin_buf.Bytes())

	if err != nil {
		fmt.Println(err)
		return
	}

	buffer := make([]byte, MESSAGE_SIZE)
	n, _, err := c.ReadFromUDP(buffer)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Reply: %s\n", string(buffer[0:n]))
}

func server(port string) {
	port = ":" + port

	s, err := net.ResolveUDPAddr("udp4", port)
	if err != nil {
		fmt.Println(err)
		return
	}

	c, err := net.ListenUDP("udp4", s)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer c.Close()
	tmp := make([]byte, MESSAGE_SIZE)
	rand.Seed(time.Now().Unix())

	for {
		//_, addr, err := c.ReadFromUDP(tmp)
		c.ReadFromUDP(tmp)
		tmpbuff := bytes.NewBuffer(tmp)
		tmpstruct := new(Message)

		// creates a decoder object
		gobobj := gob.NewDecoder(tmpbuff)
		// decodes buffer and unmarshals it into a Message struct
		gobobj.Decode(tmpstruct)

		msg := *tmpstruct

		go handleInc(msg)
	}

}

func random(min, max int) int {
	return rand.Intn(max-min) + min
}
