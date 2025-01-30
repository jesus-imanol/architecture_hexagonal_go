package controllers

import (
    "demo/src/users/application"
    "demo/src/users/domain/entities"
    "demo/src/users/infraestructure/utils"
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/dgrijalva/jwt-go"
    "time"
    "os"
)

type RegisterUserController struct {
    registerUserUseCase *application.RegisterUserUseCase
}

func NewRegisterUserController(registerUseCase *application.RegisterUserUseCase) *RegisterUserController {
    return &RegisterUserController{registerUserUseCase: registerUseCase}
}

func GenerateJWT(user entities.User) (string, error) {
    var mySigningKey = []byte(os.Getenv("JWT_SECRET"))

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "userName": user.Name,
        "exp":      time.Now().Add(time.Hour * 72).Unix(),
    })

    tokenString, err := token.SignedString(mySigningKey)
    if err != nil {
        return "", err
    }

    return tokenString, nil
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
    registerToUser, err2 := ruc.registerUserUseCase.Execute(user.Name, user.LastName, passwordHashed, user.Role)
    if err2 != nil {
        g.JSON(http.StatusInternalServerError, gin.H{"error": err2.Error()})
        return
    }
    token, err3 := GenerateJWT(user)
    if err3 != nil {
        g.JSON(http.StatusInternalServerError, gin.H{"error": err3.Error()})
        return
    }

    // Agregar el token en la cabecera de la respuesta
    g.Header("Authorization", "Bearer " + token)

    response := gin.H{
        "data": gin.H{
            "type": "users",
            "id":   registerToUser.Id,
            "attributes": gin.H{
                "name":     registerToUser.Name,
                "lastname": registerToUser.LastName,
                "role":     registerToUser.Role,
            },
        },
    }
    g.JSON(http.StatusCreated, response)
}
