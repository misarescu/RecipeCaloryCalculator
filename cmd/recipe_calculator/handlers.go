package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"recipe-calculator.com/internal/server"
)

func GetRecipe(w http.ResponseWriter, r *http.Request) error {
	recipe := Recipe{}
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&recipe); err != nil {
		return fmt.Errorf("error unmarshalling json from body")
	}

	server.WriteJSON(w, http.StatusOK, recipe)

	return nil
}
