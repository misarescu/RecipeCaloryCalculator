package main

import (
	"recipe-calculator.com/internal/server"
)

type MessageResponse struct {
	Message string `json:"message_response"`
}

func main() {
	api, _ := server.New("localhost", "8080")

	api.AddRouteHandler("GET /recipe", GetRecipe, server.DefaultErrorHandler)

	api.Run()
}
