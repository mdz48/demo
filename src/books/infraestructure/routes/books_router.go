package routes

import (
	booksControllers "demo/src/books/infraestructure/controllers"
	"github.com/gin-gonic/gin"
)

type BookRouter struct {
	engine         *gin.Engine
	bookController *booksControllers.BookController
}

func NewBookRouter(engine *gin.Engine, bookController *booksControllers.BookController) *BookRouter {
	return &BookRouter{
		engine:         engine,
		bookController: bookController,
	}
}

func (r *BookRouter) SetupRoutes() {
	books := r.engine.Group("/books")
	{
		books.POST("", r.bookController.Create)
	}
}

func (r *BookRouter) Run() error {
	return r.engine.Run()
}
