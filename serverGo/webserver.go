package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", helloServer)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal("Listen and Serve", err.Error())
	}
}

func helloServer(w http.ResponseWriter, req *http.Request) {
	fmt.Println("inside Hello Server Handler")
	fmt.Fprintf(w, "Hello,"+req.URL.Path[1:])
}
