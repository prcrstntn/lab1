package main

import (
	"net/http"
	"fmt"
	"log"
)

const PORT = ":8795"


func main() {
	http.HandleFunc("/time", TimeHandler)
	fmt.Printf("Starting server at port %s", PORT)
	err := http.ListenAndServe(PORT, nil)
	if err != nil {
	  log.Fatal(err)
	}
}