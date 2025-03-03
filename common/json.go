package common

import (
	"encoding/json"
	"net/http"
)

func WriteJSON(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(data)
}

func ReadJSON(r *http.Request, data interface{}) error {
	return json.NewDecoder(r.Body).Decode(data)
}

func WriteError(w http.ResponseWriter, code int, message string) {
	WriteJSON(w, code, map[string]string{"error": message})
}