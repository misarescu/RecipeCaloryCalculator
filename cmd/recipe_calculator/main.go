package main

import (
	"encoding/json"
	"net/http"

	"recipe-calculator.com/internal/server"
)

type response struct {
	Message string
}

func main() {
	api, _ := server.New("localhost", "8080")

	api.AddRouteHandlerFancy("GET /message", func(w http.ResponseWriter, r *http.Request) error {
		w.Header().Set("Contenty-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		res := response{Message: "all is ok"}
		err := json.NewEncoder(w).Encode(res)
		return err
	}, nil)

	api.Run()
}
