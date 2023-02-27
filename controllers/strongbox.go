package controllers

import "net/http"

func StrongboxHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		NewStrongboxHandler(w, r)
	case http.MethodGet:
		RetrieveStrongboxHandler(w, r)
	case http.MethodDelete:
		DeleteStrongboxHandler(w, r)
	default:
		w.Header().Set("Allow", "GET, POST, DELETE")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}
