package entities


type UpdateProduct struct {
	Name string `json:"name"`
	Price float32 `json:"price"`
	Deleted bool `json:"deleted"`
	Descripcion string `json:"descripcion"`
	Stock int32 `json:"stock"`
}