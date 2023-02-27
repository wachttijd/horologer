package controllers

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"horologer/cryptocode"
	"horologer/database"
	"horologer/models"
	"log"
	"net/http"
	"strconv"
)

type NewRequestData struct {
	AvailableAfter int64
	Text           string
}

type NewResponse struct {
	RetrieveKey string `json:"retrieve_key"`
}

func ExtractNewRequestData(r *http.Request) (NewRequestData, error) {
	err := r.ParseForm()

	if err != nil {
		return NewRequestData{}, err
	}

	AvailableAfter := r.FormValue("available_after")
	Text := r.FormValue("text")

	if AvailableAfter == "" || Text == "" {
		return NewRequestData{}, errors.New("missing required field(s)")
	}

	AvailableAfterInt, err := strconv.ParseInt(AvailableAfter, 10, 64)

	if err != nil {
		return NewRequestData{}, errors.New("invalid 'available_after' field format")
	}

	return NewRequestData{
		AvailableAfter: AvailableAfterInt,
		Text:           Text,
	}, nil
}

func NewStrongboxHandler(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Content-Type") != "application/x-www-form-urlencoded" {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		return
	}

	extractedData, err := ExtractNewRequestData(r)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newInternalKey, err := cryptocode.SecureRandomBytes(32)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	newUserKey, err := cryptocode.SecureRandomBytes(32)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	newGeneralId, err := cryptocode.ID()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	newAvailableAfter, err := cryptocode.EncryptAES(
		newUserKey,
		cryptocode.Int64ToBytes(extractedData.AvailableAfter),
	)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	newData, err := cryptocode.EncryptAES(
		newInternalKey,
		[]byte(extractedData.Text),
	)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	dataHash := sha256.Sum256([]byte(extractedData.Text))

	newDataIntegrity, err := cryptocode.EncryptAES(
		newInternalKey,
		dataHash[:],
	)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	newEncryptedInternalKey, err := cryptocode.EncryptAES(
		newUserKey,
		newInternalKey,
	)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	err = database.AddNewStrongbox(models.Strongbox{
		GeneralId:      newGeneralId,
		AvailableAfter: newAvailableAfter,
		DecryptionKey:  newEncryptedInternalKey,
		Integrity:      newDataIntegrity,
		Data:           newData,
	})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(NewResponse{
		RetrieveKey: newGeneralId + base64.URLEncoding.EncodeToString(newUserKey),
	})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
}
