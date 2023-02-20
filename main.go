package main

import (
	"horologer/controllers"
	"log"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/time", controllers.TimeHandler)

	server := &http.Server{
		Addr:         "127.0.0.1:8999",
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Println("Starting server...")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
