package mapi

import (
	"encoding/json"
	"github.com/valyala/fasthttp"
)

// Params stores dynamic path parameters
type Params map[string]string

// Context wraps fasthttp.RequestCtx and adds helper methods
type Context struct {
	*fasthttp.RequestCtx
	Params Params
	Next   func()
	index  int
}

// OK sends a 200 OK response with a JSON body
func (c *Context) OK(data interface{}) {
	c.JSON(fasthttp.StatusOK, data)
}

// NotFound sends a 404 Not Found response
func (c *Context) NotFound(message string) {
	c.String(fasthttp.StatusNotFound, message)
}

// BadRequest sends a 400 Bad Request response
func (c *Context) BadRequest(message string) {
	c.String(fasthttp.StatusBadRequest, message)
}

// InternalServerError sends a 500 Internal Server Error response
func (c *Context) InternalServerError(message string) {
	c.String(fasthttp.StatusInternalServerError, message)
}

// JSON sends a JSON response with the specified status code
func (c *Context) JSON(statusCode int, data interface{}) {
	c.SetContentType("application/json")
	c.SetStatusCode(statusCode)
	if err := json.NewEncoder(c).Encode(data); err != nil {
		c.Error("Failed to encode JSON", fasthttp.StatusInternalServerError)
	}
}

// BindJSON binds the request body JSON to the given destination struct
func (c *Context) BindJSON(dest interface{}) error {
	if err := json.Unmarshal(c.PostBody(), dest); err != nil {
		c.InternalServerError("Invalid JSON format")
		return err
	}
	return nil
}

// QueryParam retrieves a query parameter by name
func (c *Context) QueryParam(key string) string {
	return string(c.QueryArgs().Peek(key))
}

// SetHeader sets a response header
func (c *Context) SetHeader(key, value string) {
	c.Response.Header.Set(key, value)
}

// GetHeader retrieves a request header
func (c *Context) GetHeader(key string) string {
	return string(c.Request.Header.Peek(key))
}

// String sends a plain text response
func (c *Context) String(statusCode int, text string) {
	c.SetContentType("text/plain")
	c.SetStatusCode(statusCode)
	c.SetBodyString(text)
}

// Redirect sends a redirect response
func (c *Context) Redirect(url string, statusCode int) {
	c.Response.Header.Set("Location", url)
	c.SetStatusCode(statusCode)
}

// JSONP sends a JSONP response
func (c *Context) JSONP(callback string, data interface{}) {
	c.SetContentType("application/javascript")
	c.SetBodyString(callback + "(")
	c.JSON(fasthttp.StatusOK, data)
	c.SetBodyString(")")
}
