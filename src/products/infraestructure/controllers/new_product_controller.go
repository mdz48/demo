package controllers

import (
	"demo/src/products/application"
	"demo/src/products/domain"
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
	p := pc.createProductUseCase.Run(product)
	c.JSON(http.StatusOK, gin.H{"data": p})
}
