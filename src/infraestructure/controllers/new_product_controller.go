package controllers

import (
	"demo/src/application"
	"demo/src/domain"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	createProductUseCase *application.CreateProductUseCase
}

func NewProductController(createUseCase *application.CreateProductUseCase) *ProductController {
	return &ProductController{createProductUseCase: createUseCase}
}

func (pc *ProductController) Create(c *gin.Context) {
    var product domain.Product
    if err := c.ShouldBindJSON(&product); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    fmt.Println("Creando producto", product) // Ahora imprime el producto después de la vinculación
    pc.createProductUseCase.Run(product)
    c.JSON(http.StatusOK, gin.H{"message": "Product created successfully"})
}