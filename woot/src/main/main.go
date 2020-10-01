package main

import (
	"fmt"
	"net"
	"time"
)

const MAIN_HOSTNAME = "woot_testrun_1"
const PORT = "3002"

var ME Contact

func main() {
	inits()
	time.Sleep(5 * time.Second)

	if ME.Address != getMainIP() {
		fmt.Println("Im not main")
		sendInit()
	} else {
		fmt.Println("Im main")
	}
	server(PORT)
}

/*
	Getting the local IP and checking if its the same IP
	as MAIN_HOSTNAME which will be the node that all connect
	to.
	Then if it isn't the main it will ping the main to make it
	self known.

*/
func inits() {
	// Makes sure everything is up before trying to do stuff
	localIPs, _ := net.InterfaceAddrs()

	// The IP comes with a /16 that needs to be removed
	ME.Address = localIPs[1].String()[:(len(localIPs[1].String()) - 3)]
	fmt.Println("IP " + ME.Address)
	ME.ID = randomKadId()
}

func getMainIP() string {
	ips, err := net.LookupIP(MAIN_HOSTNAME)
	if err != nil {
		return "error"
	}
	return ips[0].String()
}

func sendInit() {
	time.Sleep(5 * time.Second)
	go client(MAIN_HOSTNAME, PORT, Message{CMD: "PING", Contact: ME})
}
