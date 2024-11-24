package mapi

import (
	"github.com/valyala/fasthttp"
	"os"
	"path/filepath"
)

func ServeStatic(directory string) HandlerFunc {
	return func(ctx *Context) {
		// Construct the full file path
		filePath := filepath.Join(directory, string(ctx.Path()))
		// Check if the file exists
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			ctx.NotFound("File not found")
			return
		}

		// Serve the file content
		data, err := os.ReadFile(filePath)
		if err != nil {
			ctx.InternalServerError("Failed to read file")
			return
		}

		// Set appropriate content type based on file extension
		ext := filepath.Ext(filePath)
		switch ext {
		case ".html":
			ctx.SetContentType("text/html")
		case ".css":
			ctx.SetContentType("text/css")
		case ".js":
			ctx.SetContentType("application/javascript")
		case ".png":
			ctx.SetContentType("image/png")
		case ".jpg", ".jpeg":
			ctx.SetContentType("image/jpeg")
			// Add other content types as needed
		}

		// Write the file content to the response body
		ctx.SetStatusCode(fasthttp.StatusOK)
		ctx.SetBody(data)
	}
}
