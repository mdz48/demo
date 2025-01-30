package controllers

import (
	"demo/src/books/application"
	"demo/src/books/domain"
	"net/http"
	"fmt"

	"github.com/gin-gonic/gin"
)

type BookController struct {
	createBookUseCase *application.CreateBookUseCase
}

func NewBookController(createUseCase *application.CreateBookUseCase) *BookController {
	return &BookController{createBookUseCase: createUseCase}
}

func (pc *BookController) Create(c *gin.Context) {
	var book domain.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	p, err := pc.createBookUseCase.Run(book)
	if err != nil {
		if err.Error() == fmt.Sprintf("el autor con ID %d no existe", book.Author) {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear el libro"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": p})
}