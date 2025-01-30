package controllers

import (
	"demo/src/users/application"
	"net/http"
	"github.com/gin-gonic/gin"
	"fmt"
)

type DeleteUsersController struct {
	deleteUserUseCase *application.DeleteUseCase
}

func NewDeleteUserController(deleteUserUseCase *application.DeleteUseCase) *DeleteUsersController {
	return &DeleteUsersController{deleteUserUseCase: deleteUserUseCase}
}

func (pc *DeleteUsersController) Delete(c *gin.Context) {
	idStr := c.Param("id")
	var id int32
	if _, err := fmt.Sscanf(idStr, "%d", &id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	result, err := pc.deleteUserUseCase.Run(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar el usuario"})
		return
	}

	if result == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
		return
	}

	c.Status(http.StatusNoContent)
}