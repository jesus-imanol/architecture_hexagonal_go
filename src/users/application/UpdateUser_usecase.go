package application

import "demo/src/users/domain/repositories"

type UpdateUserUseCase struct {
	db repositories.IUser
}

func NewUpdateUserUseCase(db repositories.IUser) *UpdateUserUseCase {
    return &UpdateUserUseCase{db: db}
}

func (uuc *UpdateUserUseCase) Execute(id int32, name string, lastname string, password string,role int32) error {
	err := uuc.db.Update(id, name, lastname, password,role)
	if err != nil {
        return err
    }
    return nil
}