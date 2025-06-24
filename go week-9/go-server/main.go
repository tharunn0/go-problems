package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/http2"
)

func main() {
	// handler func
	mux := http.NewServeMux()

	mux.HandleFunc("/hai", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "hai from handler")
	})
	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "hello from handler")
	})

	// load cert and key
	cert := "cert.pem"
	key := "key.pem"

	// configure tls
	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
	}

	// create and serve custom server
	port := 9090
	server := &http.Server{
		Addr:      fmt.Sprintf(":%d", port),
		TLSConfig: tlsConfig,
		Handler:   mux,
	}

	http2.ConfigureServer(server, &http2.Server{})

	// err := http.ListenAndServeTLS(fmt.Sprintf(":%d", port), cert, key, mux)

	err := server.ListenAndServeTLS(cert, key)
	if err != nil {
		log.Fatalln("Error starting server :", err)
	}

}
