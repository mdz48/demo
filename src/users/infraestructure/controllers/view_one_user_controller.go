package controllers

import (
	"demo/src/users/application"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ViewOneUserController struct {
	viewOneUserUseCase *application.ViewOneUserUseCase
}

func NewViewOneUserController(viewOneUserUseCase *application.ViewOneUserUseCase) *ViewOneUserController {
	return &ViewOneUserController{viewOneUserUseCase: viewOneUserUseCase}
}

func (uc *ViewOneUserController) ViewOne(c *gin.Context) {
	idStr := c.Param("id")
	var id int32
	if _, err := fmt.Sscanf(idStr, "%d", &id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	result, err := uc.viewOneUserUseCase.Run(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener el usuario"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": result})
}