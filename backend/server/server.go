package server

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Server struct {
	Address string
	Port    int
}

func New(address string, port int) *Server {
	s := &Server{
		Address: address,
		Port:    port,
	}
	return s
}

func (s *Server) Run() error {
	ctx := context.Background()

	mux := http.NewServeMux()
	mux.HandleFunc("/", s.baseHandler)

	hs := &http.Server{
		Addr:           fmt.Sprintf("%s:%d", s.Address, s.Port),
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	done := make(chan struct{}, 1)

	go func() {
		<-sigc
		log.Print("signal received, shutting down server")

		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		if err := hs.Shutdown(ctx); err != nil {
			log.Printf("error shutting down HTTP server: %s", err)
		}
		done <- struct{}{}
	}()

	log.Printf("Starting server on %s:%d", s.Address, s.Port)
	if err := hs.ListenAndServe(); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			log.Printf("HTTP server error: %s", err)
		}
	}

	<-done

	return nil
}
