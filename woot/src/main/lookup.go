package main

type Message struct {
	CMD  string
	Data string
	IP   string
	ID   byte
}

func lookup(MyId, BootstrapId byte) {
	sendPackage(Message{CMD: "lookup", ID: MyId})
}
