package server

import (
	"context"
	"log"
	"net/http"
	"time"
)

type APIServer struct {
	httpServer *http.Server
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(time.Second * 5)

	w.Write([]byte("Hello, World!"))
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("About Page"))
}

func NewAPIServer() *APIServer {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", HomeHandler)
	mux.HandleFunc("GET /about", AboutHandler)

	return &APIServer{
		httpServer: &http.Server{
			Addr:    ":8080",
			Handler: mux,
		},
	}
}

func (s *APIServer) Start() error {
	log.Printf("Server started at %s", s.httpServer.Addr)

	return s.httpServer.ListenAndServe()
}

func (s *APIServer) Stop(ctx context.Context) error {
	log.Printf("Server stopping...")

	return s.httpServer.Shutdown(ctx)
}
