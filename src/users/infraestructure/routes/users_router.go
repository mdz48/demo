package routes

import (
    usersControllers "demo/src/users/infraestructure/controllers"
    "github.com/gin-gonic/gin"
)

type UserRouter struct {
    engine         *gin.Engine
    userController *usersControllers.UserController
    viewUserController *usersControllers.ViewUsersController
    deleteUserController *usersControllers.DeleteUsersController
    updateUserController *usersControllers.UpdateUserController
    loginUserController *usersControllers.LoginUserController
    viewOneUserController *usersControllers.ViewOneUserController

}

func NewUserRouter(engine *gin.Engine,userController *usersControllers.UserController, viewUserController *usersControllers.ViewUsersController, deleteUserController *usersControllers.DeleteUsersController, updateUserController *usersControllers.UpdateUserController, loginUserController *usersControllers.LoginUserController, viewOneUserController *usersControllers.ViewOneUserController ) *UserRouter {
    return &UserRouter{
        engine:         engine,
        userController: userController,
        viewUserController: viewUserController,
        deleteUserController: deleteUserController,
        updateUserController: updateUserController,
        loginUserController: loginUserController,
        viewOneUserController: viewOneUserController,
    }
}

func (r *UserRouter) SetupRoutes() {
    users := r.engine.Group("/users")
    {
        users.POST("/", r.userController.Create)
        users.GET("/", r.viewUserController.View)
        users.GET("/:id", r.viewOneUserController.ViewOne)
        users.DELETE("/:id", r.deleteUserController.Delete)
        users.PUT("/:id", r.updateUserController.Update)
    
        login := users.Group("/login")
        {
            login.POST("", r.loginUserController.Login)
        }
    }
}

func (r *UserRouter) Run() error {
    return r.engine.Run()
}