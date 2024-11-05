package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/vpapidokha/go-graceful-shutdown/internal/server"
)

func main() {
	if err := runApp(); err != nil {
		log.Fatal(err)
	}
}

func runApp() error {
	var err error
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	wg := sync.WaitGroup{}
	wg.Add(1)

	apiServer := server.NewAPIServer()

	go func() error {
		defer wg.Done()

		err = func(apiServer *server.APIServer) error {
			return apiServer.Start()
		}(apiServer)
		if err != nil {
			return err
		}

		return nil
	}()

	<-signals
	log.Printf("App shutting down gracefully.")

	const shutdownTimeout = 5 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	if err := apiServer.Stop(ctx); err != nil {
		log.Printf("Server stopping err: %v", err)
	}

	return nil
}
