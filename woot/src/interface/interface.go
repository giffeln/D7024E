package main

func main() {
	client("localhost", "3002", Message{CMD: "PING", Data: "hejsan"})
}
