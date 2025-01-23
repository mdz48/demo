package controllers

import (
	"demo/src/application"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DeleteProductController struct {
	deleteProductUseCase *application.DeleteUseCase
}

func NewDeleteProductController(deleteUseCase *application.DeleteUseCase) *DeleteProductController {
	return &DeleteProductController{deleteProductUseCase: deleteUseCase}
}

func (pc *DeleteProductController) Delete(c *gin.Context) {
	idStr := c.Param("id")
	var id int32
	if _, err := fmt.Sscanf(idStr, "%d", &id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	result, err := pc.deleteProductUseCase.Run(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar el producto"})
		return
	}

	if result == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Producto no encontrado"})
		return
	}

	c.Status(http.StatusNoContent)
}
