package application

import (
	_"demo/src/products/domain/entities"
	"demo/src/products/domain/repositories"
)

type UpdateProduct struct {
    db repositories.IProduct
}

func NewUpdateProduct(db repositories.IProduct) *UpdateProduct {
    return &UpdateProduct{db: db}
}

func (up *UpdateProduct) Execute(id int32,name string, price float32, descripcion string, stock int32) error {
    err := up.db.Update(id, name,price,descripcion,stock)
    if err != nil {
        return err
    }
    return nil
}
