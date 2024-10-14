package submitserver

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"time"
)

type HttpServerOptions struct {
	address  string
	port     int
	certFile string
	keyFile  string
}

type SparkSubmitServer struct {
	// TODO: Add fields for the server configuration
	httpOpts   *HttpServerOptions
	httpServer *http.Server
}

func DefaultServerOptions() *HttpServerOptions {
	return &HttpServerOptions{address: "localhost", port: 15115}
}

func NewSparkSubmitServer(opts *HttpServerOptions) *SparkSubmitServer {
	return &SparkSubmitServer{}
}

func (s *SparkSubmitServer) Start() {
	cert, err := tls.LoadX509KeyPair(s.httpOpts.certFile, s.httpOpts.keyFile)
	if err != nil {
		panic(err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/health", s.Health)
	mux.HandleFunc("/submit", s.Submit)
	mux.HandleFunc("/state", s.AppState)
	mux.HandleFunc("/status", s.AppStatus)
	mux.HandleFunc("/logs", s.AppLogs)
	mux.HandleFunc("/describe", s.DescribeApp)
	mux.HandleFunc("/delete", s.DeleteApp)

	server := &http.Server{
		Addr: fmt.Sprintf(":%d", s.httpOpts.port),
		TLSConfig: &tls.Config{
			Certificates: []tls.Certificate{cert},
		},
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      mux,
	}

	if err := server.ListenAndServeTLS("", ""); err != nil {
		panic(err)
	}
	s.httpServer = server
}

func (s *SparkSubmitServer) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.httpServer.Shutdown(ctx); err != nil {
		log.Printf("HTTP server shutdown error: %v", err)
	}
	log.Println("HTTP server gracefully stopped")
}

func main() {
	opts := DefaultServerOptions()
	server := NewSparkSubmitServer(opts)
	server.Start()
	server.Stop()
}

// pod [initContainers, webhook, submitserver, controller]
// rest -> submitserver.svc -> kubeapi crd create -> classic spark operator flow
