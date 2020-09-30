package main

import "net"

func main() {
	server("3002")
}

func init() {
	ips, err := net.LookupIP("google.com")
	if err != nil {
		return
	}

}
