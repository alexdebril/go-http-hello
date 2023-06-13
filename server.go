package main

import (
	"fmt"
	"time"

	"net/http"
)

type server struct {
	logger
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	logContext := map[string]string{"method": r.Method}
	switch r.Method {
	case http.MethodGet:
		s.get(w)
		s.logEvent("success", logContext)
	default:
		s.unsupported(w)
		s.logEvent("unsupported method", logContext)
	}
}

const helloMessage = `{"message": "hello world!", "date" : "%s"}`

func (s *server) get(w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
	payload := fmt.Sprintf(helloMessage, time.Now().Format(time.DateTime))
	_, _ = w.Write([]byte(payload))
}

const unsupportedMessage = `{"message": "unsupported HTTP method. Please send a valid GET HTTP request to get a well-formed hello"}`

func (s *server) unsupported(w http.ResponseWriter) {
	w.WriteHeader(http.StatusMethodNotAllowed)
	_, _ = w.Write([]byte(unsupportedMessage))
}
