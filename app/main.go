package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gabelula/messagebird_api/app/routes"
	"github.com/urfave/negroni"
)

func init() {
	os.Setenv("LOGGING_LEVEL", "1")
}

func main() {
	if os.Getenv("MESSAGEBIRD_API_KEY") == "" {
		log.Print("MESSAGEBIRD_API_KEY environment variable is empty. Set it up with your development API key.")
		os.Exit(1)
	}

	host := os.Getenv("HOST")
	if host == "" {
		host = ":3000"
	}

	n := negroni.Classic() // Includes some default middlewares
	n.UseHandler(routes.API())

	log.Printf("Started : Listening on: %s", host)

	err := http.ListenAndServe(host, n)

	log.Printf("Down : %v", err)
}
