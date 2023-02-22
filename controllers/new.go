package controllers

import (
	"errors"
	"net/http"
	"strconv"
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

func ExtractNewRequestData(r *http.Request) (NewRequestData, error) {
	err := r.ParseForm()

	if err != nil {
		return NewRequestData{}, err
	}

	AvailableAfter := r.FormValue("available_after")
	AvailableUntil := r.FormValue("available_until")
	Text := r.FormValue("text")

	if AvailableAfter == "" || AvailableUntil == "" || Text == "" {
		return NewRequestData{}, errors.New("missing required field(s)")
	}

	AvailableAfterInt, err := strconv.ParseInt(AvailableAfter, 10, 64)

	if err != nil {
		return NewRequestData{}, errors.New("invalid 'available_after' field format")
	}

	AvailableUntilInt, err := strconv.ParseInt(AvailableUntil, 10, 64)

	if err != nil {
		return NewRequestData{}, errors.New("invalid 'available_until' field format")
	}

	return NewRequestData{
		AvailableAfter: AvailableAfterInt,
		AvailableUntil: AvailableUntilInt,
		Text: Text,
	}, nil
}

func NewHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	if r.Header.Get("Content-Type") != "application/x-www-form-urlencoded" {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		return
	}

	// ExtractedData, err := ExtractNewRequestData(r)

	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }

	// ...
}
