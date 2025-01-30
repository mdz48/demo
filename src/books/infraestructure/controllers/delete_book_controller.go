package controllers

import(
	"demo/src/books/application"
	"net/http"
	"fmt"

	"github.com/gin-gonic/gin"
)

type DeleteBookController struct {
	deleteBookUseCase *application.DeleteUseCase
}

func NewDeleteBookController(deleteUseCase *application.DeleteUseCase) *DeleteBookController {
	return &DeleteBookController{deleteBookUseCase: deleteUseCase}
}

func (dc *DeleteBookController) Delete(c *gin.Context) {
	idStr := c.Param("id")
	var id int32
	if _, err := fmt.Sscanf(idStr, "%d", &id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	result, err := dc.deleteBookUseCase.Run(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar el libro"})
		return
	}
	if result == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Libro no encontrado"})
		return
	}
	c.Status(http.StatusNoContent)
}