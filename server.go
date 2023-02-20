package main

import (
	"net/http"
	"fmt"
	"log"
)

func main() {
	fmt.Printf("Starting server at port :8795")
	err := http.ListenAndServe(":8795", nil)
	if err != nil {
	  log.Fatal(err)
	}
}