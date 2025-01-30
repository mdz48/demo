package controllers

import (
	"demo/src/books/application"
	"demo/src/books/domain"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UpdateBookController struct {
	updateBookUseCase *application.UpdateUseCase
}

func NewUpdateBookController(updateUseCase *application.UpdateUseCase) *UpdateBookController {
	return &UpdateBookController{updateBookUseCase: updateUseCase}
}

func (uc *UpdateBookController) Update(c *gin.Context) {
	idStr := c.Param("id")
	var id int32
	if _, err := fmt.Sscanf(idStr, "%d", &id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var book domain.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book"})
		return
	}

	result, err := uc.updateBookUseCase.Run(id, book)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar el libro"})
		return
	}
	c.JSON(http.StatusOK, result)
}