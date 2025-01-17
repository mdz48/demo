package main

import (
	"demo/src/infraestructure"
	"github.com/gin-gonic/gin"
)

func main() {
	r  := gin.Default()
	r.GET("/products", func(c *gin.Context){
		
	})

	r.Run()
}