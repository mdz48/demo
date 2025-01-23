package controllers

import (
	"demo/src/application"
	// "demo/src/domain"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ViewProductsController struct {
	getAllProductsUseCase *application.ViewUseCase
}

func NewViewProductsController(getAllProductsUseCase *application.ViewUseCase) *ViewProductsController {
	return &ViewProductsController{getAllProductsUseCase: getAllProductsUseCase}
}

func (pc *ViewProductsController) View(c *gin.Context) {
	fmt.Println("Obteniendo productos")
	products := pc.getAllProductsUseCase.Run()
	c.JSON(http.StatusOK, gin.H{ "data" : products})
}