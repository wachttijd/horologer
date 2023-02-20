package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type TimeResponse struct {
	Timestamp int64 `json:"timestamp"`
}

func TimeHandler(w http.ResponseWriter, r *http.Request) {
	response := TimeResponse{
		Timestamp: time.Now().Unix(),
	}

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
	}
}
