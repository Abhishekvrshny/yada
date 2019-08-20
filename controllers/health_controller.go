package controllers

import "net/http"

// Health is the HTTP handler for health request
func Health(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}
