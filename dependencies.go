package main

import (
	"demo/src/core"
	"demo/src/products/application"
	"demo/src/products/infraestructure"
	controllers2 "demo/src/products/infraestructure/controllers"
	"demo/src/products/infraestructure/routes"
	usersUseCases "demo/src/users/application"
	usersInfraestructure "demo/src/users/infraestructure"
	usersControllers "demo/src/users/infraestructure/controllers"
	usersRoutes "demo/src/users/infraestructure/routes"

	"github.com/gin-gonic/gin"

	booksUseCases "demo/src/books/application"
	booksInfraestructure "demo/src/books/infraestructure"
	booksControllers "demo/src/books/infraestructure/controllers"
	booksRoutes "demo/src/books/infraestructure/routes"
)

type Dependencies struct {
	engine *gin.Engine
}

func NewDependencies() *Dependencies {
	return &Dependencies{
		engine: gin.Default(),
	}
}

func (d *Dependencies) Run() error {
	database := core.NewDatabase()

	// Products setup
	db := infraestructure.NewMySQL(database.Conn)
	createUseCase := application.NewCreateUseCase(db)
	productController := controllers2.NewProductController(createUseCase)
	viewUseCase := application.NewUseCaseCreate(db)
	viewProductsController := controllers2.NewViewProductsController(viewUseCase)
	updateUseCase := application.NewUseCaseUpdate(db)
	updateProductController := controllers2.NewUpdateProductController(updateUseCase)
	deleteUseCase := application.NewUseCaseDelete(db)
	deleteProductController := controllers2.NewDeleteProductController(deleteUseCase)

	// Configurar rutas de productos
	productsRouter := routes.NewRouter(d.engine, productController, viewProductsController, updateProductController, deleteProductController)
	productsRouter.SetupRoutes()

	// Users setup
	usersDatabase := usersInfraestructure.NewMySQL(database.Conn)
	createUserUseCase := usersUseCases.NewCreateUseCase(usersDatabase)
	userController := usersControllers.NewUserController(createUserUseCase)
	viewUserUseCase := usersUseCases.NewUseCaseView(usersDatabase)
	viewUserController := usersControllers.NewViewUsersController(viewUserUseCase)
	deleteUserUseCase := usersUseCases.NewUseCaseDelete(usersDatabase)
	deleteUserController := usersControllers.NewDeleteUserController(deleteUserUseCase)
	updateUserUseCase := usersUseCases.NewUseCaseUpdate(usersDatabase)
	updateUserController := usersControllers.NewUpdateUserController(updateUserUseCase)
	loginUserUseCase := usersUseCases.NewUseCaseLogin(usersDatabase)
	loginUserController := usersControllers.NewLoginUserController(loginUserUseCase)

	// Configurar rutas de usuarios
	usersRouter := usersRoutes.NewUserRouter(d.engine, userController, viewUserController, deleteUserController, updateUserController, loginUserController)
	usersRouter.SetupRoutes()

	// Books setup
	dbBooks := booksInfraestructure.NewMySQL(database.Conn)
	viewBooksUseCase := booksUseCases.NewUseCaseViewBooks(dbBooks)
	viewBooksController := booksControllers.NewViewBooksController(viewBooksUseCase)
	createUseCaseBooks := booksUseCases.NewCreateUseCase(dbBooks)
	bookController := booksControllers.NewBookController(createUseCaseBooks)
	deleteUseCaseBooks := booksUseCases.NewUseCaseDelete(dbBooks)
	deleteBookController := booksControllers.NewDeleteBookController(deleteUseCaseBooks)
	updateBookUseCase := booksUseCases.NewUseCaseUpdate(dbBooks)
	updateBookController := booksControllers.NewUpdateBookController(updateBookUseCase)
    viewBooksByAuthorUseCase := booksUseCases.NewViewBooksByAuthorUseCase(dbBooks)
    viewBooksByAuthorController := booksControllers.NewViewAuthorByUserController(viewBooksByAuthorUseCase)
	addFavoriteBookUseCase := booksUseCases.NewAddFavoriteBookUseCase(dbBooks)
	addFavoriteBookController := booksControllers.NewAddFavoriteBookController(addFavoriteBookUseCase)

	// Configurar rutas de libros
	booksRouter := booksRoutes.NewBookRouter(d.engine, bookController, deleteBookController, updateBookController, viewBooksController, viewBooksByAuthorController, addFavoriteBookController)
	booksRouter.SetupRoutes()

	return d.engine.Run()
}
