package main

import (
	"github.com/gin-contrib/cors"
)

func main() {
	dependencies := NewDependencies()

	// Configurar CORS
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true                                             // Permitir todos los orígenes
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"} // Permitir métodos
	config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization"} // Permitir encabezados

	dependencies.engine.Use(cors.New(config)) // Agregar middleware de CORS

	dependencies.Run()
}
