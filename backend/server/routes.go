package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

// Response describes the payload sent back to a client querying this webapp.
type Response struct {
	Method     string            `json:"method"`
	URL        string            `json:"url"`
	Proto      string            `json:"proto"`
	Headers    map[string]string `json:"headers"`
	Host       string            `json:"host"`
	RequestURI string            `json:"request_uri"`
}

func (s *Server) baseHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("new request: %+v\n", r)

	headers := make(map[string]string)
	for k, v := range r.Header {
		headers[k] = strings.Join(v, ", ")
	}

	response := Response{
		Method:     r.Method,
		URL:        r.URL.String(),
		Proto:      r.Proto,
		Headers:    headers,
		Host:       r.Host,
		RequestURI: r.RequestURI,
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(response); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "ERROR: %s\n", err)
		return
	}
}

func (s *Server) largeHandler(w http.ResponseWriter, r *http.Request) {
	b := make([]byte, 8192)
	for i := range b {
		b[i] = 'A'
	}

	w.Header().Add("X-Large-Header", string(b))
	fmt.Fprintf(w, "very large\n")
}
