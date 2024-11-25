
# minimal_api

**minimal_api** es un microframework ligero y de alto rendimiento escrito en Go, diseñado para facilitar la creación de APIs minimalistas. Su diseño está inspirado en frameworks como Flask y FastAPI, pero construido para aprovechar la velocidad de `fasthttp`.

Este proyecto combina simplicidad y eficiencia, ofreciendo una experiencia sencilla para desarrolladores que buscan rapidez y flexibilidad en sus aplicaciones.

---

## Características principales

- **Ultra rápido**: Construido sobre `fasthttp`, ofrece un rendimiento superior para manejar solicitudes HTTP.
- **API limpia y sencilla**: Define rutas, middlewares y manejadores de manera intuitiva.
- **Middlewares**: Soporte para middlewares globales y específicos para cada ruta.
- **Apagado controlado**: Manejo adecuado de señales del sistema para detener el servidor de manera segura.
- **Código modular**: Separación de responsabilidades entre el servidor y el enrutador para facilitar la escalabilidad.

---

## Instalación

1. **Clonar el repositorio**:

   ```bash
   git clone https://github.com/pol-cova/minimal_api.git
   cd minimal_api
   ```

2. **Instalar las dependencias**:

   Asegúrate de tener Go 1.18 o superior instalado y ejecuta:

   ```bash
   go mod tidy
   ```

3. **Ejecutar el servidor**:

   ```bash
   go run main.go
   ```

---

## Ejemplo básico

Este ejemplo muestra cómo usar `minimal_api` para configurar un servidor simple con rutas y middlewares:

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

### Prueba el servidor

- Visita [http://localhost:8080/](http://localhost:8080/) para ver el mensaje de bienvenida.
- Visita [http://localhost:8080/protegido](http://localhost:8080/protegido) para probar una ruta con middleware específico.

---

## Cómo detener el servidor

El servidor está diseñado para apagarse de forma segura cuando recibe una señal de interrupción. 

1. Inicia el servidor normalmente:

   ```bash
   go run main.go
   ```

2. Detén el servidor presionando `Ctrl+C` en la terminal. Verás un mensaje como:

   ```
   Shutting down server...
   Server gracefully stopped
   ```

---

## Estructura del proyecto

```plaintext
minimal_api/
├── main.go         // Archivo principal para iniciar el servidor
├── server.go       // Lógica principal del servidor
├── router.go       // Gestión de rutas y middlewares
├── go.mod          // Dependencias del proyecto
├── go.sum          // Hashes de las dependencias
```

---

## Contribuciones

¡Eres bienvenido a contribuir! Aquí tienes algunas formas de hacerlo:

1. **Reporta errores**: Abre un [issue](https://github.com/pol-cova/minimal_api/issues) para problemas o preguntas.
2. **Mejora el código**: Envía un *pull request* con nuevas características o correcciones.
3. **Documenta**: Ayuda a mejorar la documentación con ejemplos adicionales o guías.

---

## Licencia

Este proyecto está bajo la licencia MIT. Consulta el archivo [LICENSE](LICENSE) para más información.

---

Si necesitas ayuda, no dudes en crear un issue en el repositorio o contactarme directamente. ¡Gracias por tu interés en minimal_api!
