package utils

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

// this is for caching, using same instance of validator for all requests
var Validate = validator.New()

func ParseJson(r *http.Request, payloadVar any) error {
	if r.Body == nil {
		return fmt.Errorf("empty body")
	}

	return json.NewDecoder(r.Body).Decode(&payloadVar)
}

func WriteJson(w http.ResponseWriter, status int, output any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(&output)
}

func WriteError(w http.ResponseWriter, status int, err error) error {
	return WriteJson(w, status, map[string]string{"error": err.Error()})
}
