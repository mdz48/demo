package main

import (
	application2 "demo/src/products/application"
	"demo/src/products/core"
	"demo/src/products/infraestructure"
	controllers2 "demo/src/products/infraestructure/controllers"
	"demo/src/products/infraestructure/routes"
	// "github.com/gin-gonic/gin"
)

func main() {
	// Crear instancias
	database := core.NewDatabase()
	db := infraestructure.NewMySQL(database.Conn)
	createUseCase := application2.NewCreateUseCase(db)
	productController := controllers2.NewProductController(createUseCase)
	// Cada controlador necesita su propio caso de uso, es necesario instanciar cada vez?
	viewUseCase := application2.NewUseCaseCreate(db)
	viewProductsController := controllers2.NewViewProductsController(viewUseCase)
	updateUseCase := application2.NewUseCaseUpdate(db)
	updateProductController := controllers2.NewUpdateProductController(updateUseCase)

	deleteUseCase := application2.NewUseCaseDelete(db)
	deleteProductController := controllers2.NewDeleteProductController(deleteUseCase)

	// Configurar rutas
	r := routes.NewRouter(productController, viewProductsController, updateProductController, deleteProductController)
	r.SetupRoutes()

	// En donde irá mi archivo de rutas?

	// Iniciar servidor
	r.Run()
}
