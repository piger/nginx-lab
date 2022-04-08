// This program is a simple web application used to simulate a backend service exposed through nginx.
package main

import (
	"flag"
	"log"

	"github.com/piger/nginx-lab/server"
)

var (
	listenAddress = flag.String("addr", "0.0.0.0", "Address to bind to")
	listenPort    = flag.Int("port", 4444, "Port to listen on")
)

func run() error {
	flag.Parse()

	srv := server.New(*listenAddress, *listenPort)
	if err := srv.Run(); err != nil {
		return err
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatalf("error: %s", err)
	}
}
