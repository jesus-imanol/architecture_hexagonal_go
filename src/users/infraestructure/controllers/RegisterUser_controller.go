package controllers

import (
	"demo/src/users/application"
	"demo/src/users/domain/entities"

	"net/http"

	"github.com/gin-gonic/gin"
	"demo/src/users/infraestructure/utils"
)

type RegisterUserController struct {
	registerUserUseCase *application.RegisterUserUseCase
}

func NewRegisterUserController(registerUseCase *application.RegisterUserUseCase) *RegisterUserController {
    return &RegisterUserController{registerUserUseCase: registerUseCase}
}


func (ruc *RegisterUserController) RegisterUser(g *gin.Context) {
	var user entities.User
	if err := g.ShouldBindJSON(&user); err != nil {
        g.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error()})
        return
    }
	passwordHashed, err := utils.HashPassword(user.Password)
	if err!= nil {
        g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
	registerToUser, err2 := ruc.registerUserUseCase.Execute(user.Name,user.LastName, passwordHashed, user.Role)
	if err2 != nil {
        g.JSON(http.StatusInternalServerError, gin.H{"error": err2.Error()})
        return
    }
	response := gin.H{
        "data": gin.H{
            "type": "users",
            "id":   registerToUser.Id,
            "attributes": gin.H{
                "name":  registerToUser.Name,
                "lastname": registerToUser.LastName,
                "role": registerToUser.Role,
            },
        },
    }
    g.JSON(http.StatusCreated, response)
}

