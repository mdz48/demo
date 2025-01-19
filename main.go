package main

import (
	"demo/src/application"
	"demo/src/infraestructure"
	"demo/src/infraestructure/controllers"
	"demo/src/infraestructure/routes"
	// "github.com/gin-gonic/gin"
)

func main() {
    // Crear instancias
    db := infraestructure.NewMySQL()
    createUseCase := application.NewCreateUseCase(db)
    productController := controllers.NewProductController(createUseCase)
    // Cada controlador necesita su propio caso de uso, es necesario instanciar cada vez?
    viewUseCase := application.NewUseCaseCreate(db)
    viewProductsController := controllers.NewViewProductsController(viewUseCase)
    updateUseCase := application.NewUseCaseUpdate(db)
    updateProductController := controllers.NewUpdateProductController(updateUseCase)
    
    deleteUseCase := application.NewUseCaseDelete(db)
    deleteProductController := controllers.NewDeleteProductController(deleteUseCase)


    // Configurar rutas
    r := routes.NewRouter(productController, viewProductsController, updateProductController, deleteProductController)
    r.SetupRoutes()
    
    // En donde irá mi archivo de rutas?

    // Iniciar servidor
    r.Run()
}