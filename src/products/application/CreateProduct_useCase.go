package application

import (
	"demo/src/products/domain/entities"
	"demo/src/products/domain/repositories"

)

type CreateProduct struct {
    db repositories.IProduct
}

func NewCreateProduct(db repositories.IProduct) *CreateProduct {
    return &CreateProduct{db: db}
}

func (cp *CreateProduct) Execute(name string, price float32, descripcion string, stock int32, user_id int32) (*entities.Product, error) {
    product := entities.NewProduct(name, price, stock, descripcion, user_id)
    err := cp.db.Save(product)
    if err != nil {
        return nil, err
    }
    return product, nil
}
