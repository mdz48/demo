package controllers

import (
	"demo/src/books/application"
	"demo/src/books/domain"
	// "fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type AddFavoriteBookController struct {
    addFavoriteBookUseCase *application.AddFavoriteBookUseCase
}

func NewAddFavoriteBookController(addFavoriteBookUseCase *application.AddFavoriteBookUseCase) *AddFavoriteBookController {
    return &AddFavoriteBookController{addFavoriteBookUseCase: addFavoriteBookUseCase}
}

func (c *AddFavoriteBookController) Add(g *gin.Context) {
    var request domain.FavoriteBookRequest
    if err := g.ShouldBindJSON(&request); err != nil {
        g.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
        return
    }

    err := c.addFavoriteBookUseCase.Run(request.UserId, request.BookId)
    if err != nil {
        if strings.Contains(err.Error(), "ya está en favoritos") {
            g.JSON(http.StatusConflict, gin.H{"error": err.Error()})
            return
        }
        g.JSON(http.StatusInternalServerError, gin.H{"error": "Error al agregar libro favorito"})
        return
    }

    g.JSON(http.StatusOK, gin.H{"message": "Libro agregado a favoritos"})
}