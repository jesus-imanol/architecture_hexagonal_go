package application

import (
	"demo/src/products/domain/entities"
	"demo/src/products/domain/repositories"
	"fmt"
)

type ListProduct struct {
    db repositories.IProduct
}

func NewListProduct(db repositories.IProduct) *ListProduct {
    return &ListProduct{db: db}
}

func (lp *ListProduct) Execute() ([]*entities.Product, error) {
    products, err := lp.db.GetAll()
	fmt.Println("products", products)
    if err != nil {
        return nil, err
    }
    return products, nil
}
