package main

import (
    // "github.com/gin-gonic/gin"
    "github.com/gin-contrib/cors"
)

func main() {
    dependencies := NewDependencies()
    
    // Configurar CORS
    config := cors.DefaultConfig()
    config.AllowAllOrigins = true // Permitir todos los orígenes
    // O puedes especificar orígenes permitidos
    // config.AllowOrigins = []string{"http://localhost:4200"}

    dependencies.engine.Use(cors.New(config)) // Agregar middleware de CORS

    dependencies.Run()
}