package controllers

import (
	"demo/src/books/application"
	"net/http"
	"github.com/gin-gonic/gin"
)

type ViewBooksController struct {
	viewBooksUseCase *application.ViewBooksUseCase
}

func NewViewBooksController(viewBooksUseCase *application.ViewBooksUseCase) *ViewBooksController {
	return &ViewBooksController{viewBooksUseCase: viewBooksUseCase}
}

func (vc *ViewBooksController) View(c *gin.Context) {
	books, err := vc.viewBooksUseCase.Run()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los libros"})
		return
	}
	c.JSON(http.StatusOK, books)
}