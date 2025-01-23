package controllers

import (
	"demo/src/application"
	"demo/src/domain"
	"net/http"
	"github.com/gin-gonic/gin"
	"fmt"
)

type UpdateProductController struct {
	updateProductUseCase *application.UpdateUseCase
}

func NewUpdateProductController(updateUseCase *application.UpdateUseCase) *UpdateProductController {
	return &UpdateProductController{updateProductUseCase: updateUseCase}
}

func (pc *UpdateProductController) Update(c *gin.Context) {
	idStr := c.Param("id")
	var id int32
	if _, err := fmt.Sscanf(idStr, "%d", &id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var product domain.Product
	if err := c.BindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product"})
		return
	}

	result, err := pc.updateProductUseCase.Run(id, product)
	if err != nil {
        if err.Error() == "product not found" {
            c.JSON(http.StatusNotFound, gin.H{"error": "Producto no encontrado"})
        } else {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar el producto"})
        }
        return
    }

	c.JSON(http.StatusOK, gin.H{"data": result})
}