package application

import (
	_"demo/src/products/domain/entities"
	"demo/src/products/domain/repositories"
)
type DeleteProduct struct {
	db repositories.IProduct
}
func NewDeleteProduct(db repositories.IProduct) *DeleteProduct {
	return &DeleteProduct{db: db}
}
func (dp *DeleteProduct) Execute(id int32) error {
	err := dp.db.Delete(id)
	if err != nil {
        return err
    }
	return nil
}