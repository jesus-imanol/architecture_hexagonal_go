package repositories

import "demo/src/products/domain/entities"

type IProduct interface {
	Save(product *entities.Product) error
	GetAll() ([]*entities.Product, error)
	Delete(id int32) error
	Update(id int32,name string, price float32, descripcion string, stock int32) error
}
