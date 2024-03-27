package server

import "net/http"

type Server struct {
	listenAddr string
	listenPort string
	router     *http.ServeMux
}

type ApiHandler func(http.ResponseWriter, *http.Request) error
type ErrorHandler func(w http.ResponseWriter, err error)
