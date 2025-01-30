package repositories

import "demo/src/users/domain/entities"

type IUser interface {
	Register(user *entities.User) error
	Update(id int32, name string, lastname string, password string, role int32) error
	GetAll() ([]*entities.User, error)
	Delete(id int32) error
} 