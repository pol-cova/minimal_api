package mapi

import (
	"strings"

	"github.com/valyala/fasthttp"
)

type HandlerFunc func(ctx *Context)

type Node struct {
	part        string
	handler     HandlerFunc
	middlewares []HandlerFunc
	children    map[string]*Node
}

type Router struct {
	root        *Node
	middlewares []HandlerFunc
}

// NewRouter initializes a new Router
func NewRouter() *Router {
	return &Router{
		root:        &Node{children: make(map[string]*Node)},
		middlewares: []HandlerFunc{},
	}
}

// Use registers a global middleware
func (r *Router) Use(middleware HandlerFunc) {
	r.middlewares = append(r.middlewares, middleware)
}

// AddRoute registers a new route with optional middlewares
func (r *Router) AddRoute(method, path string, handler HandlerFunc, middlewares ...HandlerFunc) {
	parts := strings.Split(path, "/")[1:] // Split path into parts (excluding the leading "/")
	node := r.root

	for _, part := range parts {
		if _, exists := node.children[part]; !exists {
			node.children[part] = &Node{
				part:     part,
				children: make(map[string]*Node),
			}
		}
		node = node.children[part]
	}

	node.handler = handler
	node.middlewares = middlewares
}

// GET registers a GET route
func (r *Router) GET(path string, handler HandlerFunc, middlewares ...HandlerFunc) {
	r.AddRoute("GET", path, handler, middlewares...)
}

// POST registers a POST route
func (r *Router) POST(path string, handler HandlerFunc, middlewares ...HandlerFunc) {
	r.AddRoute("POST", path, handler, middlewares...)
}

// PUT registers a PUT route
func (r *Router) PUT(path string, handler HandlerFunc, middlewares ...HandlerFunc) {
	r.AddRoute("PUT", path, handler, middlewares...)
}

// DELETE registers a DELETE route
func (r *Router) DELETE(path string, handler HandlerFunc, middlewares ...HandlerFunc) {
	r.AddRoute("DELETE", path, handler, middlewares...)
}

// Handler processes incoming HTTP requests
func (r *Router) Handler(ctx *fasthttp.RequestCtx) {
	path := string(ctx.Path())

	// Split the path into segments and search in the radix tree
	parts := strings.Split(path, "/")[1:]

	// Start matching the path in the radix tree
	node := r.root
	params := make(Params)

	for _, part := range parts {
		if child, exists := node.children[part]; exists {
			node = child
		} else if handler, exists := node.children[":*"]; exists {
			// Wildcard route, matches any segment
			node = handler
			params["*"] = part
		} else {
			// Route not found
			ctx.Response.SetStatusCode(fasthttp.StatusNotFound)
			ctx.Response.SetBodyString("404 Not Found")
			return
		}
	}

	// Execute global middlewares, route-specific middlewares, and handler
	c := &Context{RequestCtx: ctx, Params: params}
	executeChain(c, append(r.middlewares, append(node.middlewares, node.handler)...))
}

// executeChain runs middlewares and handler in order
func executeChain(c *Context, chain []HandlerFunc) {
	for _, fn := range chain {
		fn(c)
	}
}
