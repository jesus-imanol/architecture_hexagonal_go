package entities


type Product struct {
	ID int32 `json:"product_id"`
	Name string `json:"name"`
	Price float32 `json:"price"`
	Deleted bool `json:"deleted"`
	Descripcion string `json:"descripcion"`
	Stock int32 `json:"stock"`
}

func NewProduct(name string, price float32,stock int32,descripcion string) *Product{
	return &Product{Name: name, Price: price,Stock: stock ,Descripcion: descripcion, Deleted: false}
}

func (p *Product) GetName() string {
	return p.Name
}

func (p *Product) SetName(name string) {
	p.Name = name
}