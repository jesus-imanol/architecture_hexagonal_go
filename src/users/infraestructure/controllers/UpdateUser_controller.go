package controllers

import (
	"demo/src/users/application"
	"demo/src/users/domain/entities"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UpdateUserController struct {
	updateUserUseCase *application.UpdateUserUseCase
}

func NewUUpdateUserController(updateUseCase * application.UpdateUserUseCase) *UpdateUserController {
	return &UpdateUserController{updateUserUseCase: updateUseCase}
}

func (uuc *UpdateUserController) UpdateUser(g *gin.Context) {
	var user entities.User
    idParam := g.Param("id")
    id, err := strconv.ParseInt(idParam, 10, 32)
    if err!= nil {
        g.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
        return
    }

    if err := g.ShouldBindJSON(&user); err!= nil {
        g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
	idCheck := int32(id)

    err = uuc.updateUserUseCase.Execute(idCheck, user.Name, user.LastName, user.Password, user.Role)
    if err!= nil {
        g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
	response := gin.H {
		"data": gin.H{
            "type": "users",
            "id":   idCheck,
            "attributes": gin.H{
                "name":      user.Name,
                "last_name": user.LastName,
                "password":  user.Password,
                "role":       user.Role,
            },
        },
    }
    g.JSON(http.StatusOK, response)
	}

