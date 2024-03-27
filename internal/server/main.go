package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

var shutdown os.Signal = syscall.SIGUSR1

type Server struct {
	listenAddr string
	listenPort string
	router     *http.ServeMux
}

type ApiHandler func(http.ResponseWriter, *http.Request) error
type ErrorHandler func(error)

func MakeApiHandler(f ApiHandler, e ErrorHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			if e != nil {
				e(err) // handle the error here
			}
		}
	}
}

func New(addr, port string) (*Server, error) {
	return &Server{listenAddr: addr, listenPort: port, router: http.NewServeMux()}, nil
}

func (s *Server) AddRouteHandler(apiRoute string, f ApiHandler, e ErrorHandler) {
	s.router.HandleFunc(apiRoute, MakeApiHandler(f, e))
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
