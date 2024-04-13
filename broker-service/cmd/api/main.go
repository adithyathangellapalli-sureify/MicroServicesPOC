package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = "8082"

type Config struct{}

func main() {
	app := Config{}

	log.Printf("Broker Service starting on port %s\n", port)
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: app.routes(),
	}

	error := server.ListenAndServe()

	if error != nil {
		log.Panic(error)
	}

}
