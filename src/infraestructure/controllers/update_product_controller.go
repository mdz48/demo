package controllers

import (
	"demo/src/application"
	"demo/src/domain"
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)

type UpdateProductController struct {
	updateProductUseCase *application.UpdateUseCase
}

func NewUpdateProductController(updateUseCase *application.UpdateUseCase) *UpdateProductController {
	return &UpdateProductController{updateProductUseCase: updateUseCase}
}

func (pc *UpdateProductController) Update(c *gin.Context) {
	var product domain.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("Actualizando producto", product) // Ahora imprime el producto después de la vinculación
	pc.updateProductUseCase.Run(product)
	c.JSON(http.StatusOK, gin.H{"message": "Product updated successfully"})
}