package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

func main() {
	var PORT string
	if PORT = os.Getenv("PORT"); PORT == "" {
		PORT = "3001"
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var message string = r.URL.Path[1:]
		message = strings.ToLower(message)
		if message == "hello" {
			fmt.Fprintf(w, "world!")
		} else {
			fmt.Fprintf(w, "Go to /Hello")
		}
		//fmt.Fprintf(w, "Hello World from path: %s\n", r.URL.Path)
	})
	http.ListenAndServe(":"+PORT, nil)
}

func waitForInput() {
	//var test string
}
