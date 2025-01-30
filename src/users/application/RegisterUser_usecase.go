package application

import (
	"demo/src/users/domain/entities"
	"demo/src/users/domain/repositories"
)

type RegisterUserUseCase struct {
	db repositories.IUser
}

func NewRegisterUserUseCase(db repositories.IUser) *RegisterUserUseCase {
    return &RegisterUserUseCase{db: db}
}

func (ru *RegisterUserUseCase) Execute(name string, lastname string, password string,role int32) (*entities.User,error){
	user := entities.NewUser(name, lastname,password,role)
	err := ru.db.Register(user)
	if err!= nil {
        return nil,err
    }
    return user,nil
}