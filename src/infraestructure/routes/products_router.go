package routes

import (
    "github.com/gin-gonic/gin"
    "demo/src/infraestructure/controllers"
)

type Router struct {
    engine *gin.Engine
    productController      *controllers.ProductController
    viewProductsController *controllers.ViewProductsController
    updateProductController *controllers.UpdateProductController
    deleteProductController *controllers.DeleteProductController
}

func NewRouter(productController *controllers.ProductController, viewProductsController *controllers.ViewProductsController, updateProductController *controllers.UpdateProductController, deleteProductController *controllers.DeleteProductController) *Router {
    return &Router{
        engine:                gin.Default(),
        productController:     productController,
        viewProductsController: viewProductsController,
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
        products.PUT("", r.updateProductController.Update)
        products.DELETE("/:id", r.deleteProductController.Delete)
    }
}

func (r *Router) Run() error {
    return r.engine.Run()
}