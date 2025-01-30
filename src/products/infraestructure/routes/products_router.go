package routes

import (
	controllers2 "demo/src/products/infraestructure/controllers"
	"github.com/gin-gonic/gin"
)

type Router struct {
	engine                  *gin.Engine
	productController       *controllers2.ProductController
	viewProductsController  *controllers2.ViewProductsController
	updateProductController *controllers2.UpdateProductController
	deleteProductController *controllers2.DeleteProductController
}

func NewRouter(engine *gin.Engine,productController *controllers2.ProductController, viewProductsController *controllers2.ViewProductsController, updateProductController *controllers2.UpdateProductController, deleteProductController *controllers2.DeleteProductController) *Router {
	return &Router{
		engine:                  gin.Default(),
		productController:       productController,
		viewProductsController:  viewProductsController,
		updateProductController: updateProductController,
		deleteProductController: deleteProductController,
	}
}

func (r *Router) SetupRoutes() {
	// Grupo de rutas para productos
	products := r.engine.Group("/products")
	{
		products.POST("", r.productController.Create)
		products.GET("", r.viewProductsController.View)
		products.PUT("/:id", r.updateProductController.Update)
		products.DELETE("/:id", r.deleteProductController.Delete)
	}
}

func (r *Router) Run() error {
	return r.engine.Run()
}
