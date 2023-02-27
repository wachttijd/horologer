package controllers

import (
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"horologer/cryptocode"
	"horologer/database"
	"log"
	"net/http"
)

func DeleteStrongboxHandler(w http.ResponseWriter, r *http.Request) {
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

	dataIntegrity, err := cryptocode.DecryptAES(
		decryptionKey,
		box.Integrity,
	)

	if err != nil {
		http.Error(w, "decryption failed, the given key might be invalid", http.StatusForbidden)
		log.Println(err)
		return
	}

	decryptedDataHash := sha256.Sum256(decryptedText)

	if hex.EncodeToString(decryptedDataHash[:]) != hex.EncodeToString(dataIntegrity) {
		http.Error(w, "decryption failed, the given key might be invalid", http.StatusForbidden)
		log.Println(err)
		return
	}

	if err := database.DeleteStrongbox(generalId); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
