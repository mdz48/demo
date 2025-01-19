package controllers

import (
	"demo/src/application"
	"net/http"
	"github.com/gin-gonic/gin"
    "fmt"
)

type DeleteProductController struct {
	deleteProductUseCase *application.DeleteUseCase
}

func NewDeleteProductController(deleteUseCase *application.DeleteUseCase) *DeleteProductController {
	return &DeleteProductController{deleteProductUseCase: deleteUseCase}
}

func (pc *DeleteProductController) Delete(c *gin.Context) {
    idStr := c.Param("id") // Obtener el ID de la URL
    var id int32
    // Convertir el ID de string a int32
    if _, err := fmt.Sscanf(idStr, "%d", &id); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }
    pc.deleteProductUseCase.Run(id)
    c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}