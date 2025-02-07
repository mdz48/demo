package controllers 

import (
	"demo/src/books/application"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ViewFavoritesController struct {
	getFavoritesUseCase *application.GetFavoritesUseCase
}

func NewViewFavoritesController(getFavoritesUseCase *application.GetFavoritesUseCase) *ViewFavoritesController {
	return &ViewFavoritesController{getFavoritesUseCase: getFavoritesUseCase}
}

func (c *ViewFavoritesController) View(g *gin.Context) {
    userId := g.Param("userId")
    userIdInt, err := strconv.ParseInt(userId, 10, 32)
    if err != nil {
        g.JSON(http.StatusBadRequest, gin.H{"error": "ID de usuario inválido"})
        return
    }

    favorites, err := c.getFavoritesUseCase.Run(int32(userIdInt))
    if err != nil {
        g.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener libros favoritos"})
        return
    }

    g.JSON(http.StatusOK, favorites)
}