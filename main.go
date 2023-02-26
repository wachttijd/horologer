package main

import (
	"horologer/conf"
	"horologer/controllers"
	"log"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/time", controllers.TimeHandler)
	mux.HandleFunc("/strongbox", controllers.StrongboxHandler)

	config, err := conf.RetrieveConfig()

	if err != nil {
		log.Fatal(err)
	}

	server := &http.Server{
		Addr:         config.Server.Address,
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Printf("Running server on %s", config.Server.Address)

	err = server.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}
