package main

import (
	"net/http"
	"strconv"

	"recipe-calculator.com/internal/server"
)

func main() {
	api, _ := server.New("localhost", "8080")

	api.AddRouteHandler("GET /message", func(w http.ResponseWriter, r *http.Request) error {
		res := server.Response{Message: "all is ok"}
		err := server.WriteJSON(w, http.StatusOK, res)
		return err
	}, nil)

	api.AddRouteHandler("GET /message/{id}", func(w http.ResponseWriter, r *http.Request) error {
		strId := r.PathValue("id")
		if _, err := strconv.Atoi(strId); err != nil {
			return err
		}

		res := server.Response{Message: "id is string and ok"}

		err := server.WriteJSON(w, http.StatusOK, res)

		return err
	}, server.DefaultErrorHandler)

	api.Run()
}
