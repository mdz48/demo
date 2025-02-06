package routes

import (
	booksControllers "demo/src/books/infraestructure/controllers"

	"github.com/gin-gonic/gin"
)

type BookRouter struct {
	engine                      *gin.Engine
	bookController              *booksControllers.BookController
	deleteBookController        *booksControllers.DeleteBookController
	updateBookController        *booksControllers.UpdateBookController
	viewBooksController         *booksControllers.ViewBooksController
	viewBooksByAuthorController *booksControllers.ViewBooksByAuthorController
	addFavoriteBookController   *booksControllers.AddFavoriteBookController
}

func NewBookRouter(engine *gin.Engine, bookController *booksControllers.BookController, deleteBookController *booksControllers.DeleteBookController, updateBookController *booksControllers.UpdateBookController, viewBooksController *booksControllers.ViewBooksController, viewBooksByAuthorController *booksControllers.ViewBooksByAuthorController, addFavoriteBookController *booksControllers.AddFavoriteBookController) *BookRouter {
	return &BookRouter{
		engine:                      engine,
		bookController:              bookController,
		deleteBookController:        deleteBookController,
		updateBookController:        updateBookController,
		viewBooksController:         viewBooksController,
		viewBooksByAuthorController: viewBooksByAuthorController,
		addFavoriteBookController:   addFavoriteBookController,
	}
}

func (r *BookRouter) SetupRoutes() {
	books := r.engine.Group("/books")
	{
		books.GET("", r.viewBooksController.View)
		books.POST("", r.bookController.Create)
		books.DELETE("/:id", r.deleteBookController.Delete)
		books.PUT("/:id", r.updateBookController.Update)
		books.GET("/author/:authorId", r.viewBooksByAuthorController.View)

		favorites := books.Group("/favorites")
		{
			favorites.POST("", r.addFavoriteBookController.Add)
		}
	}
}

func (r *BookRouter) Run() error {
	return r.engine.Run()
}
