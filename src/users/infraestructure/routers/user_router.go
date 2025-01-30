package routers
import (
	"demo/src/users/infraestructure/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine, registerController *controllers.RegisterUserController, updateController *controllers.UpdateUserController, listUserController *controllers.ListUserController, deleteController *controllers.DeleteUserController) {
	v1 := r.Group("/v1/users")
	{
		v1.POST("/", registerController.RegisterUser)
		v1.PUT("/:id", updateController.UpdateUser)
		v1.GET("/", listUserController.GetAllUsers)
		v1.DELETE("/:id", deleteController.DeleteUser)
	}
}
	