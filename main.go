package main

import (
	"log"
	"os"

	"net/http"
)

const defaultPort = ":8080"

func main() {
	l := logger{os.Stdout}
	s := &server{
		logger: l,
	}

	envPort := os.Getenv("HTTP_PORT")
	port := defaultPort
	if envPort != "" {
		port = envPort
	}
	l.logEvent("listening", map[string]string{"port": port})
	err := http.ListenAndServe(port, s)
	if err != nil {
		log.Fatalln(err)
	}
}
