package controllers

import (
    "demo/src/users/application"
    "demo/src/users/domain"
    "net/http"
    "github.com/gin-gonic/gin"
)

type UserController struct {
    createUserUseCase *application.CreateUserUseCase
}

func NewUserController(createUseCase *application.CreateUserUseCase) *UserController {
    return &UserController{createUserUseCase: createUseCase}
}

func (pc *UserController) Create(c *gin.Context) {
    var user domain.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    u := pc.createUserUseCase.Run(user)
    c.JSON(http.StatusOK, gin.H{"data": u})
}