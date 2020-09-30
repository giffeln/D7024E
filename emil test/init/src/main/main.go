package main

import (
	"fmt"
	"net"
	"time"
)

const MAIN_HOSTNAME = "init_testrun_1"

func main() {
	go inits()
	server("3002")
}

func inits() {
	time.Sleep(3 * time.Second)
	ips, err := net.LookupIP(MAIN_HOSTNAME)
	if err != nil {
		return
	}
	localIPs, _ := net.InterfaceAddrs()
	localIP := localIPs[1].String()[:(len(localIPs[1].String()) - 3)]
	time.Sleep(3 * time.Second)

	if ips[0].String() == localIP {
		fmt.Println("Yah")
		return
	}
	fmt.Println("Nah")
	time.Sleep(5 * time.Second)
	client(MAIN_HOSTNAME, "3002", Message{CMD: "PING"})
}
