package controllers

import(
	"demo/src/books/application"
	"net/http"
	"fmt"
	"strings"
	"github.com/gin-gonic/gin"
)

type DeleteFavoriteController struct {
	deleteFavoriteUseCase *application.DeleteFavoriteUseCase
}

func NewDeleteFavoriteController(deleteFavoriteUseCase *application.DeleteFavoriteUseCase) *DeleteFavoriteController {
	return &DeleteFavoriteController{deleteFavoriteUseCase: deleteFavoriteUseCase}
}

func (dc *DeleteFavoriteController) Delete(c *gin.Context) {
	userIdStr := c.Param("userId")
	bookIdStr := c.Param("bookId")
	var userId int32
	var bookId int32
	if _, err := fmt.Sscanf(userIdStr, "%d", &userId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	if _, err := fmt.Sscanf(bookIdStr, "%d", &bookId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	err := dc.deleteFavoriteUseCase.Run(userId, bookId)
    if err != nil {
        if strings.Contains(err.Error(), "no se encontró") {
            c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar el libro de favoritos"})
        return
    }
    c.Status(http.StatusNoContent)
}