package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func requestLogger(logger *log.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			next.ServeHTTP(w, r)
			logger.Printf("Served request: %s %s", r.Method, r.URL.Path)
		})
	}
}

func initializeLogger(logFile string) (*log.Logger, error) {
	if logFile != "" {
		file, err := os.OpenFile("linko.access.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0o644)
		if err != nil {
			return nil, fmt.Errorf("failed to open or create log file: %v", err)
		}
		multiWriter := io.MultiWriter(os.Stderr, file)
		return log.New(multiWriter, "", log.LstdFlags), nil
	}
	return log.New(os.Stderr, "", log.LstdFlags), nil
}
