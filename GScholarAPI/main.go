package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	// Gorilla mux package - HTTP Request multiplexer helps in request routing and dispatching to handlers
	router := GScholarAPIRouter()

	// Setting default port
	port := "4747"
	if os.Getenv("PORT")!="" {
		port = os.Getenv("PORT")
	}
	port = ":" + port

	log.Fatal(http.ListenAndServe(port, router))
}