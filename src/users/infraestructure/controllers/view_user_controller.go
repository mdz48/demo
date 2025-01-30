package controllers

import (
	"demo/src/users/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ViewUsersController struct {
	getAllUsersUseCase *application.ViewUserUseCase
}

func NewViewUsersController(getAllUsersUseCase *application.ViewUserUseCase) *ViewUsersController {
	return &ViewUsersController{getAllUsersUseCase: getAllUsersUseCase}
}

func (pc *ViewUsersController) View(c *gin.Context) {
	users, err := pc.getAllUsersUseCase.Run()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": users})
}