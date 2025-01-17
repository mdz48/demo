package infraestructure

import (
	"demo/src/domain"
	"fmt"
)

type MySQL struct{
	products []domain.Product
}

func NewMySQL() *MySQL {
	return &MySQL{
		products: []domain.Product{},
	}
}

func (m *MySQL) Save(product domain.Product) {
	// Simulacion
	m.products = append(m.products, product)
	fmt.Println("Guardando en MySQL")
}

func (m *MySQL) GetAll() []domain.Product {
	// Simulacion
	fmt.Println("Obteniendo de MySQL")
	return m.products
}
