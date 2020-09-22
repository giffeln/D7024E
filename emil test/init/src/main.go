package main

func main() {
	go server("3002")

	client("localhost", "3002", Message{ID: "PING", Data: "hejsan"})
}
