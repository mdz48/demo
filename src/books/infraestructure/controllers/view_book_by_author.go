package controllers

import (
	"demo/src/books/application"
	"net/http"
	"github.com/gin-gonic/gin"
	"fmt"
)

type ViewBooksByAuthorController struct {
	viewBooksByAuthorUseCase *application.ViewBooksByAuthorUseCase
}

func NewViewAuthorByUserController(viewBooksByUserUseCase *application.ViewBooksByAuthorUseCase) *ViewBooksByAuthorController {
	return &ViewBooksByAuthorController{viewBooksByAuthorUseCase: viewBooksByUserUseCase}
}

func (vc *ViewBooksByAuthorController) View(c *gin.Context) {
	authorId := c.Param("authorId")
    var id int32
    if _, err := fmt.Sscanf(authorId, "%d", &id); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "ID de autor inválido"})
        return
    }

    books, err := vc.viewBooksByAuthorUseCase.Run(id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los libros del autor"})
        return
    }
    c.JSON(http.StatusOK, books)
}