package mapi

import (
	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
	"time"
)

// App struct represents the main framework app
type App struct {
	router *Router
	server *fasthttp.Server
}

// NewApp initializes and returns a new App instance
func NewApp() *App {
	// Initialize the router and server
	router := NewRouter()
	server := &fasthttp.Server{
		Handler:            router.Handler,
		ReadTimeout:        10 * time.Second,
		WriteTimeout:       10 * time.Second,
		MaxRequestBodySize: 10 * 1024 * 1024,
		Name:               "minimal_api Microframework Server",
	}

	// Return the app instance
	return &App{
		router: router,
		server: server,
	}
}

// Run starts the server
func (app *App) Run(addr string) {
	logrus.Infof("Starting server on %s...", addr)
	if err := app.server.ListenAndServe(addr); err != nil {
		logrus.Fatalf("Error in starting server: %v", err)
	}
}

// UseLogger applies the logger middleware to the app
func (app *App) UseLogger() {
	app.router.Use(Logger) // Use the logging middleware globally
}

// GET registers a GET route
func (app *App) GET(path string, handler HandlerFunc, middlewares ...HandlerFunc) {
	app.router.GET(path, handler, middlewares...)
}

// POST registers a POST route
func (app *App) POST(path string, handler HandlerFunc, middlewares ...HandlerFunc) {
	app.router.POST(path, handler, middlewares...)
}

// PUT registers a PUT route
func (app *App) PUT(path string, handler HandlerFunc, middlewares ...HandlerFunc) {
	app.router.PUT(path, handler, middlewares...)
}

// DELETE registers a DELETE route
func (app *App) DELETE(path string, handler HandlerFunc, middlewares ...HandlerFunc) {
	app.router.DELETE(path, handler, middlewares...)
}
