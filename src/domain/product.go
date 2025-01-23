package domain

type Product struct {
	Id    int32
	Name  string
	Price float32
}

func NewProduct(name string, price float32) *Product {
	return &Product{Name: name, Price: price}
}

func (p *Product) GetName() string {
	return p.Name
}

func (p *Product) SetName(name string) {
	p.Name = name
}

func (p *Product) UpdateProduct(name string, price float32) {
	p.Name = name
	p.Price = price
}

func (p *Product) GetID() int32 {
	return p.Id
}