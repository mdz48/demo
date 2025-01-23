package controllers

import (
	"demo/src/application"
	"demo/src/domain"
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
	pc.updateProductUseCase.Run(product)
	c.JSON(http.StatusOK, gin.H{"data": product})
}