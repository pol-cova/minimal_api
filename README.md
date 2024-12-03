# minimal_api

**minimal_api** is a lightweight and high-performance microframework written in Go, designed to facilitate the creation of minimalist APIs. Its design is inspired by frameworks like Flask and FastAPI, but built to leverage the speed of `fasthttp`.

This project combines simplicity and efficiency, offering an easy experience for developers seeking speed and flexibility in their applications.

---

## Main Features

- **Ultra-fast**: Built on `fasthttp`, it offers superior performance for handling HTTP requests.
- **Clean and simple API**: Define routes, middlewares, and handlers intuitively.
- **Middlewares**: Support for global and route-specific middlewares.
- **Graceful shutdown**: Proper handling of system signals to safely stop the server.
- **Modular code**: Separation of responsibilities between the server and the router to facilitate scalability.

---

## Installation

1. **Clone the repository**:

   ```bash
   git clone https://github.com/pol-cova/minimal_api.git
   cd minimal_api
   ```

2. **Install dependencies**:

   Make sure you have Go 1.18 or higher installed and run:

   ```bash
   go mod tidy
   ```

3. **Run the server**:

   ```bash
   go run main.go
   ```

---

## Basic Example

This example shows how to use `minimal_api` to set up a simple server with routes and middlewares:

```go
package main

import (
   "github.com/pol-cova/minimal_api/mapi"
)

func main() {
   app := mapi.NewApp()

   app.UseLogger()

   app.GET("/api", func(c *mapi.Context) {
      c.Response.SetBodyString("Hello, from minimal API")
   })

   app.POST("/api", func(c *mapi.Context) {
      c.Response.SetBodyString("Hello, from POST")
   })

   app.Run("127.0.0.1:5000")
}
```

### Test the Server

- Visit [http://localhost:8080/api](http://localhost:8080/) to see the welcome message.
- Visit [http://localhost:8080/api](http://localhost:8080/protegido) to test a POST route.

---

## How to Stop the Server

The server is designed to shut down safely when it receives an interrupt signal.

1. Start the server normally:

   ```bash
   go run main.go
   ```

2. Stop the server by pressing `Ctrl+C` in the terminal. You will see a message like:

   ```
   Shutting down server...
   Server gracefully stopped
   ```

---

## Project Structure

```plaintext
minimal_api/
|--mapi/
   ├── app.go         // Main file to start the server
   ├── server.go      // Main server logic
   ├── router.go      // Route and middleware management
   ├── context.go     // Request and response context
   ├── response.go    // HTTP response handling
   ├── request.go     // HTTP request handling
   ├── logger.go      // Request logging middleware
   |--middleware.go   // Middleware definitions
   |--template.go     // Template for server-side HTML
   |--static.go       // Static file handling

├── go.mod            // Project dependencies
├── go.sum            // Dependency hashes
```

---

## Future Features

1. **minicli**: Command-line tool to generate projects and control the server.
2. **Documentation**: Improve documentation and examples to facilitate API usage.
3. **Performance**: Optimize server performance and efficiency.

## Contributions

You are welcome to contribute! Here are some ways to do so:

1. **Report bugs**: Open an [issue](https://github.com/pol-cova/minimal_api/issues) for problems or questions.
2. **Improve the code**: Submit a pull request with new features or fixes.
3. **Document**: Help improve the documentation with additional examples or guides.

---

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more information.

---

If you need help, feel free to create an issue in the repository or contact me directly. Thank you for your interest in minimal_api!
Contact: **polc394@gmail.com**
