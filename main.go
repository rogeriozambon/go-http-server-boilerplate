package main

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/handlers"
	"github.com/rogeriozambon/go-service-boilerplate/database"
	"github.com/rogeriozambon/go-service-boilerplate/services"
)

var (
	// serviceCertFile = os.Getenv("SERVICE_CERT_FILE")
	// serviceCertKey  = os.Getenv("SERVICE_CERT_KEY")
	serviceAddr = os.Getenv("SERVICE_ADDR")
)

func main() {
	postgres := database.NewPostgres()
	mux := http.NewServeMux()

	s := services.New(postgres.DB)
	s.Register(mux)

	if err := run(mux); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}
}

func run(mux *http.ServeMux) error {
	tlsConfig := &tls.Config{
		PreferServerCipherSuites: true,
		CurvePreferences: []tls.CurveID{
			tls.CurveP256,
			tls.X25519,
		},
		MinVersion: tls.VersionTLS12,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,
			tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,
			tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
		},
	}

	withMiddlewares := handlers.LoggingHandler(os.Stdout, mux)
	withMiddlewares = handlers.CompressHandler(withMiddlewares)
	withMiddlewares = handlers.RecoveryHandler()(withMiddlewares)

	server := &http.Server{
		Addr:         serviceAddr,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
		TLSConfig:    tlsConfig,
		Handler:      withMiddlewares,
	}

	// return server.ListenAndServeTLS(serviceCertFile, serviceCertKey)
	return server.ListenAndServe()
}
