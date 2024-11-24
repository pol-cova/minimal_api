package mapi

import (
	"github.com/valyala/fasthttp"
)

type Middleware func(ctx *Context) bool

func AuthMiddleware(ctx *Context) bool {
	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "" {
		ctx.JSON(fasthttp.StatusUnauthorized, map[string]string{"error": "Unauthorized"})
		return false // Stop further processing
	}
	return true
}
