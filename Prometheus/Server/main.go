package main 

import (
	"fmt"
	"log"
	"net/http"
)

func ping(w http.ResponseWriter, req *http.Request) {
	log.Println("Got ping! Sending pong!")
	fmt.Fprintf(w, "pong")
}

func main() {

	http.HandleFunc("/ping", ping)
	log.Println("Starting on 8090, /ping is available")
	http.ListenAndServe(":8090", nil)
}
