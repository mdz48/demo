package controllers

import (
	"demo/src/users/domain"
	"demo/src/users/application"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UpdateUserController struct {
	updateUserUseCase *application.UpdateUseCase
}

func NewUpdateUserController(updateUserUseCase *application.UpdateUseCase) *UpdateUserController {
	return &UpdateUserController{updateUserUseCase: updateUserUseCase}
}

func (uc *UpdateUserController) Update(c *gin.Context) {
	idStr := c.Param("id")
	var id int32
	if _, err := fmt.Sscanf(idStr, "%d", &id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var user domain.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user"})
		return
	}

	result, err := uc.updateUserUseCase.Run(id, user)
	if err != nil {
		if err.Error() == "user not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar el usuario"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": result})
}

