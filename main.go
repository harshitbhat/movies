package main

import (
	"log"
	"net/http"

	"harshitbhat.com/movies/logger"
)

// returns a pointer(can return nil as well)
func initialiseLogger() *logger.Logger {
	logInstance, err := logger.NewLogger("logs/movies.log")
	if err != nil {
		log.Fatalf("Failed to initialise logger %v", err)
	}

	defer logInstance.Close()
	return logInstance
}

func main() {

	logInstance := initialiseLogger()

	http.Handle("/", http.FileServer(http.Dir("public")))

	const addr = ":8080"
	err := http.ListenAndServe(addr, nil)

	if err != nil {
		// sending it to the output console (can be a file as well)
		log.Fatalf("Server failed: %v", err)
		logInstance.Error("server failed", err)
	}
}