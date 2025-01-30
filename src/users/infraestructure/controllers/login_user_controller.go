package controllers

import (
	"demo/src/users/application"
	"demo/src/users/domain"
	"net/http"
	"github.com/gin-gonic/gin"
)

type LoginUserController struct {
	loginUseCase *application.LoginUserUseCase
}

func NewLoginUserController(loginUseCase *application.LoginUserUseCase) *LoginUserController {
	return &LoginUserController{loginUseCase}
}

func (controller *LoginUserController) Login(c *gin.Context) {
    var loginReq domain.LoginRequest
    
    if err := c.ShouldBindJSON(&loginReq); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    
    response, err := controller.loginUseCase.Login(loginReq.Email, loginReq.Password)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
        return
    }
    
    c.JSON(http.StatusOK, response)
}
