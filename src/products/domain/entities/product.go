package entities


type Product struct {
	ID int32 `json:"product_id"`
	Name string `json:"name"`
	Price float32 `json:"price"`
	Deleted bool `json:"deleted"`
	Descripcion string `json:"descripcion"`
	Stock int32 `json:"stock"`
	User_id int32 `json:"user_id"`
}

func NewProduct(name string, price float32,stock int32,descripcion string, user_id int32) *Product{
	return &Product{Name: name, Price: price,Stock: stock ,Descripcion: descripcion, Deleted: false, User_id: user_id}
}

func (p *Product) GetName() string {
	return p.Name
}

func (p *Product) SetName(name string) {
	p.Name = name
}