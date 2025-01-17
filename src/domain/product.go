package domain

type Product struct {
	id    int32
	name  string
	price float32
}

func NewProduct(name string, price float32) *Product {
	return &Product{id: 1, name: name, price: price}
}

func (p *Product) GetName() string {
	return p.name
}

func (p *Product) SetName(name string) {
	p.name = name
}
