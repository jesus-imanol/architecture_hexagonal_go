package routers

import (
    "demo/src/users/infraestructure/controllers"
    "demo/src/users/infraestructure/middleware"
    "github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine, registerController *controllers.RegisterUserController, updateController *controllers.UpdateUserController, listUserController *controllers.ListUserController, deleteController *controllers.DeleteUserController) {
    v1 := r.Group("/v1/users")
    {
        v1.POST("/", registerController.RegisterUser)
    }

    v1Auth := r.Group("/v1/users")
    v1Auth.Use(middleware.AuthMiddleware())
    {
        v1Auth.PUT("/:id", updateController.UpdateUser)
        v1Auth.GET("/", listUserController.GetAllUsers)
        v1Auth.DELETE("/:id", deleteController.DeleteUser)
    }
}
