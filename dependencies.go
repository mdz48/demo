package main

import (
    "demo/src/core"
    "demo/src/products/application"
    controllers2 "demo/src/products/infraestructure/controllers"
    "demo/src/products/infraestructure/routes"
    usersUseCases "demo/src/users/application"
    usersControllers "demo/src/users/infraestructure/controllers"
    usersRoutes "demo/src/users/infraestructure/routes"
    "demo/src/products/infraestructure"
    usersInfraestructure "demo/src/users/infraestructure"
    "github.com/gin-gonic/gin"

    booksInfraestructure "demo/src/books/infraestructure"
    booksControllers "demo/src/books/infraestructure/controllers"
    booksRoutes "demo/src/books/infraestructure/routes"
    booksUseCases "demo/src/books/application"
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

    // Configurar rutas de usuarios
    usersRouter := usersRoutes.NewUserRouter(d.engine, userController, viewUserController, deleteUserController, updateUserController)
    usersRouter.SetupRoutes()

    // Books setup
    dbBooks := booksInfraestructure.NewMySQL(database.Conn)
    createUseCaseBooks := booksUseCases.NewCreateUseCase(dbBooks)
    bookController := booksControllers.NewBookController(createUseCaseBooks)
    

    // Configurar rutas de libros
    booksRouter := booksRoutes.NewBookRouter(d.engine, bookController)
    booksRouter.SetupRoutes()

    return d.engine.Run()
}