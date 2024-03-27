package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"recipe-calculator.com/internal/server"
)

type response struct {
	Message string
}

func main() {
	api, _ := server.New("localhost", "8080")

	api.AddRouteHandler("GET /message", func(w http.ResponseWriter, r *http.Request) error {
		w.Header().Set("Contenty-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		res := response{Message: "all is ok"}
		err := json.NewEncoder(w).Encode(res)
		return err
	}, nil)

	api.AddRouteHandler("GET /message/{id}", func(w http.ResponseWriter, r *http.Request) error {
		w.Header().Set("Contenty-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		strId := r.PathValue("id")
		if _, err := strconv.Atoi(strId); err != nil {
			return err
		}

		res := response{Message: "id is string and ok"}

		err := json.NewEncoder(w).Encode(res)

		return err
	}, func(err error) {
		log.Println(err.Error())
	})

	api.Run()
}
