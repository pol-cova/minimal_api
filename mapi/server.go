package mapi

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/valyala/fasthttp"
)

type Server struct {
	router      *Router
	middlewares []HandlerFunc
}

// NewServer initializes a new server with a router
func NewServer() *Server {
	return &Server{
		router:      NewRouter(),
		middlewares: []HandlerFunc{},
	}
}

// Use registers a middleware at the server level
func (s *Server) Use(middleware HandlerFunc) {
	s.middlewares = append(s.middlewares, middleware)
}

// Start starts the server and listens on the given address
func (s *Server) Start(addr string) {
	log.Printf("Server started at %s\n", addr)

	// Combine middlewares with router handler
	handler := func(ctx *fasthttp.RequestCtx) {
		c := &Context{RequestCtx: ctx}

		// Apply server-level middlewares
		for _, middleware := range s.middlewares {
			middleware(c)
		}

		// Call the router handler to handle the request
		s.router.Handler(ctx)
	}

	// Set up the server with more configurations
	server := &fasthttp.Server{
		Handler:            handler,
		ReadTimeout:        10 * time.Second,
		WriteTimeout:       10 * time.Second,
		MaxRequestBodySize: 10 * 1024 * 1024,
		Name:               "minimal_api Microframework Server",
	}

	// Graceful shutdown handling with context
	go func() {
		if err := server.ListenAndServe(addr); err != nil {
			log.Fatalf("Error starting server: %v\n", err)
		}
	}()

	// Wait for an interrupt signal to shut down gracefully
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	// Gracefully shut down the server with context and timeout
	log.Println("Shutting down server...")

	shutdownTimeout := 10 * time.Second // Set a more reasonable timeout
	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	// Attempt graceful shutdown
	if err := server.Shutdown(); err != nil {
		log.Fatalf("Error shutting down server: %v\n", err)
	}

	select {
	case <-ctx.Done():
		log.Println("Server gracefully stopped")
	case <-time.After(shutdownTimeout):
		log.Println("Graceful shutdown timed out")
	}
}

// Router returns the server's router
func (s *Server) Router() *Router {
	return s.router
}
