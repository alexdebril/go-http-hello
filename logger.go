package main

import (
	"fmt"
	"io"
	"log"
	"time"
)

type logger struct {
	writer io.Writer
}

const logFormat = "%s: %s - %s\n"

func (l *logger) logEvent(msg string, contextData map[string]string) {
	c := make([]string, 0)
	for name, value := range contextData {
		c = append(c, fmt.Sprintf("%v:%v", name, value))
	}
	logMessage := fmt.Sprintf(logFormat, time.Now().Format(time.DateTime), msg, c)
	_, err := l.writer.Write([]byte(logMessage))
	if err != nil {
		log.Printf("error writing log message: %e \n", err)
		log.Println(logMessage)
	}
}
