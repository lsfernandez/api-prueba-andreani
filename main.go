package main

import (
	"github.com/architecture-it/go-platform/web"
	"github.com/customer-experience/api-prueba/api"
)

// @title Api Prueba
// @version 1.0
// @description API Prueba - Ejercicio Pintureria
// @contact.name Leo Fernandez
// @contact.url https://developers.andreani.com
// @contact.email leofernandez@andreani.com
// @host localhost:8080
// @BasePath /api/v1
func main() {
	server := web.NewServer(web.ReadConfigFromEnv())
	server.AddMetrics()
	server.AddCorsAllOrigins()
	server.AddHealth()
	server.AddApiDocs()
	api.SetupRouter(server.GetRouter())
	server.ListenAndServe()
}
