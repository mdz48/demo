package main

import (
    "demo/src/application"
    "demo/src/infraestructure"
    "github.com/gin-gonic/gin"
)

func main() {
    // Crear instancias
    db := infraestructure.NewMySQL()
    createUseCase := application.NewCreateUseCase(db)
    productController := infraestructure.NewProductController(createUseCase)
    // Cada controlador necesita su propio caso de uso, es necesario instanciar cada vez?
    viewUseCase := application.NewUseCaseCreate(db)
    viewProductsController := infraestructure.NewViewProductsController(viewUseCase)

    // Configurar rutas
    r := gin.Default()
    r.POST("/products", productController.Create)
    r.GET("/products", viewProductsController.View)

    // En donde irá mi archivo de rutas?

    // Iniciar servidor
    r.Run()
}