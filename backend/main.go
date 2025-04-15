// This program is a simple web application used to simulate a backend service exposed through nginx.
package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"strings"
	"time"
)

type logMiddleware struct {
	inner  http.Handler
	logger *slog.Logger
}

func (l logMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now().UTC()
	l.inner.ServeHTTP(w, r)
	l.logger.Info("request", "method", r.Method, "url", r.URL.String(), "remote_addr", r.RemoteAddr, "elapsed", time.Since(start).String())
}

func withLog(h http.HandlerFunc, logger *slog.Logger) http.Handler {
	return logMiddleware{inner: h, logger: logger}
}

type echoResponse struct {
	Method     string            `json:"method"`
	URL        string            `json:"url"`
	Proto      string            `json:"proto"`
	Headers    map[string]string `json:"headers"`
	Host       string            `json:"host"`
	RequestURI string            `json:"request_uri"`
	RemoteAddr string            `json:"remote_addr"`
}

// echoHandler is an handler that returns a JSON object containing information about the received request.
func echoHandler(w http.ResponseWriter, r *http.Request) {
	headers := make(map[string]string)
	for k, v := range r.Header {
		headers[k] = strings.Join(v, ", ")
	}

	resp := echoResponse{
		Method:     r.Method,
		URL:        r.URL.String(),
		Proto:      r.Proto,
		Headers:    headers,
		Host:       r.Host,
		RequestURI: r.RequestURI,
		RemoteAddr: r.RemoteAddr,
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "ERROR: %s\n", err)
		return
	}
}

func run(addr string, logger *slog.Logger) error {
	mux := http.NewServeMux()
	mux.Handle("/", withLog(echoHandler, logger))

	srv := http.Server{
		Addr:           addr,
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20, // 1 MiB
	}

	if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	return nil
}

func main() {
	flagAddr := flag.String("addr", ":4444", "Address to listen to (e.g. 0.0.0.0:4444)")
	flag.Parse()

	logger := slog.New(slog.NewJSONHandler(os.Stderr, nil))
	slog.SetDefault(logger)

	if err := run(*flagAddr, logger); err != nil {
		log.Fatalf("error: %s", err)
	}
}
