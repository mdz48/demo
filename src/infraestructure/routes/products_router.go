package routes

import (
    "github.com/gin-gonic/gin"
    "demo/src/infraestructure/controllers"
)

type Router struct {
    engine *gin.Engine
    productController      *controllers.ProductController
    viewProductsController *controllers.ViewProductsController
}

func NewRouter(productController *controllers.ProductController, viewProductsController *controllers.ViewProductsController) *Router {
    return &Router{
        engine:                gin.Default(),
        productController:     productController,
        viewProductsController: viewProductsController,
    }
}

func (r *Router) SetupRoutes() {
    // Grupo de rutas para productos
    products := r.engine.Group("/products")
    {
        products.POST("", r.productController.Create)
        products.GET("", r.viewProductsController.View)
    }
}

func (r *Router) Run() error {
    return r.engine.Run()
}