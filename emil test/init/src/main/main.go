package main

import (
	"net"
	"fmt"
	"time"
)

const MAIN_HOSTNAME = "docker_testrun_1"

func main() {
	inits()
	server("3002")
}

func inits() {
	ips, err := net.LookupIP(MAIN_HOSTNAME)
	if err != nil {
		return
	}
	if ips[0].String() == "127.0.0.1" {
		return
	}
	time.Sleep(5 * time.Second)
	client(MAIN_HOSTNAME, "3002", Message{CMD: "PING"})
	fmt.Println(ips)
}
