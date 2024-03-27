package server

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

var shutdown os.Signal = syscall.SIGUSR1

func New(addr, port string) (*Server, error) {
	return &Server{listenAddr: addr, listenPort: port, router: http.NewServeMux()}, nil
}

func (s *Server) Run() {
	addr := s.listenAddr + ":" + s.listenPort
	server := &http.Server{Addr: addr}

	http.Handle("/", s.router) // handle all routes

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	go func() {
		log.Println("starting server on ", server.Addr)
		if err := server.ListenAndServe(); err != nil {
			log.Fatalln(err.Error())
			stop <- shutdown
		}
	}()

	signal := <-stop
	log.Println("Shutting down server ...")

	server.Shutdown(context.TODO())

	if signal == shutdown {
		os.Exit(1)
	}
}

func MakeApiHandler(f ApiHandler, e ErrorHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			if e != nil {
				e(w, err) // handle the error here
			}
		}
	}
}

func (s *Server) AddRouteHandler(apiRoute string, f ApiHandler, e ErrorHandler) {
	s.router.HandleFunc(apiRoute, MakeApiHandler(f, e))
}

func LogErrorAndSendResponse(w http.ResponseWriter, res any, code int, err error) {
	WriteJSON(w, code, res)
	log.Println(err.Error())
}

func DefaultErrorHandler(w http.ResponseWriter, err error) {
	er := struct {
		Message string `json:"message"`
	}{
		Message: "Invalid Request",
	}
	LogErrorAndSendResponse(w, er, http.StatusBadRequest, err)
}

func WriteJSON(w http.ResponseWriter, code int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	return json.NewEncoder(w).Encode(data)
}
