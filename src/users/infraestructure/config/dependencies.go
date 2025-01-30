package config

import (
	"demo/src/users/infraestructure/routers"
	"demo/src/users/application"
	"demo/src/users/infraestructure/adapters"
	"demo/src/users/infraestructure/controllers"

	"github.com/gin-gonic/gin"
)


func InitUsers(r *gin.Engine) {
	ps, err := adapters.NewMySQL()
	if err != nil {
	panic(err)
	}
	registerUseCase := application.NewRegisterUserUseCase(ps)
	registerUser_controller := controllers.NewRegisterUserController(registerUseCase)

	updateUseCase := application.NewUpdateUserUseCase(ps)
	updateUser_controller := controllers.NewUUpdateUserController(updateUseCase)


	listUserUseCase := application.NewListUserUseCase(ps)
	listUser_controller := controllers.NewListUserController(listUserUseCase)

	// DELETE USER
	deleteUserUseCase := application.NewDeleteUserUseCase(ps)
	deleteUser_controller := controllers.NewDeleteUserController(deleteUserUseCase)
	routers.UserRoutes(r,registerUser_controller, updateUser_controller, listUser_controller, deleteUser_controller)
}