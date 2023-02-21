package controllers

import (
	"net/http"
)

type NewRequestData struct {
	AvailableAfter int64
	AvailableUntil int64
	Text           string
}

type NewResponse struct {
	RetrieveKey   string `json:"retrieve_key"`
	EncryptedText string `json:"encrypted_text"`
}

// func ExtractRequestData(r *http.Request) (NewRequestData, error) {
// 	err := r.ParseForm()

// 	if err != nil {
// 		return NewRequestData{}, err
// 	}

// 	AvailableAfter := r.FormValue("available_after")
// 	AvailableUntil := r.FormValue("available_until")
// 	Text := r.FormValue("text")

// 	if AvailableAfter == "" || AvailableUntil == "" || Text == "" {
// 		return NewRequestData{}, errors.New("missing required field(s)")
// 	}

// 	// ...
// }

func NewHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	if r.Header.Get("Content-Type") != "application/x-www-form-urlencoded" {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		return
	}

	// ...
}
