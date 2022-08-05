package handlers

import (
	"log"
	"net/http"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("RootHandler", r.URL)
	w.WriteHeader(404)
}
