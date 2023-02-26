package controllers

import (
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"errors"
	"horologer/cryptocode"
	"horologer/database"
	"log"
	"net/http"
	"net/url"
	"time"
)

type RetrieveResponse struct {
	Text string `json:"text"`
}

func ParseRetrieveKey(retrieveKey string) (string, []byte, error) {
	retrieveKey, err := url.QueryUnescape(retrieveKey)

	if err != nil || len(retrieveKey) <= 20 {
		return "", nil, errors.New("invalid retrieve key")
	}

	keyBytes, err := base64.URLEncoding.DecodeString(retrieveKey[20:])

	if err != nil {
		return "", nil, errors.New("invalid retrieve key")
	}

	return retrieveKey[:20], keyBytes, nil
}

func RetrieveStrongboxHandler(w http.ResponseWriter, r *http.Request) {
	currentTimestamp := time.Now().Unix()

	retrieveKey := r.URL.Query().Get("key")

	if retrieveKey == "" {
		http.Error(w, "missing 'key' parameter", http.StatusBadRequest)
		return
	}

	generalId, userKey, err := ParseRetrieveKey(retrieveKey)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	box, err := database.GetStrongbox(generalId)

	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "record not found", http.StatusNotFound)
			return
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			return
		}
	}

	availableAfter, err := cryptocode.DecryptAES(
		userKey,
		box.AvailableAfter,
	)

	if err != nil {
		http.Error(w, "decryption failed, the given key might be invalid", http.StatusForbidden)
		log.Println(err)
		return
	}

	if cryptocode.BytesToInt64(availableAfter) >= currentTimestamp {
		http.Error(w, "time condition is not met", http.StatusForbidden)
		return
	}

	decryptionKey, err := cryptocode.DecryptAES(
		userKey,
		box.DecryptionKey,
	)

	if err != nil {
		http.Error(w, "decryption failed, the given key might be invalid", http.StatusForbidden)
		log.Println(err)
		return
	}

	decryptedText, err := cryptocode.DecryptAES(
		decryptionKey,
		box.Data,
	)

	if err != nil {
		http.Error(w, "decryption failed, the given key might be invalid", http.StatusForbidden)
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(RetrieveResponse{
		Text: string(decryptedText),
	})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
}
