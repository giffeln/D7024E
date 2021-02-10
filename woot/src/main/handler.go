package main

import (
	"encoding/hex"
	"fmt"
)

/*
	Handles incoming messages

	PING sends a ping towards an adress
	Sends a RESP when it reaches the right target

	RESP recived after successful PING

*/

func handleInc(msg Message) {
	cmd := msg.CMD
	remIP := msg.Contact.Address
	remID := hex.EncodeToString(msg.Contact.ID[:])
	fmt.Println("IP: " + remIP + "\nID: " + remID)
	switch cmd {
	case "PING":
		fmt.Println("PING recieved from " + remIP)
		client(remIP, PORT, Message{CMD: "RESP", Contact: ME})
	case "RESP":
		fmt.Println("PING " + remIP + " Successful")
	case "LOOKUP":
		fmt.Println("LOOKUP recieved")
	}
}
