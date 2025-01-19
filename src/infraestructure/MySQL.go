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
	m.products = append(m.products, product)
	fmt.Println("Guardando en MySQL")
}

func (m *MySQL) GetAll() []domain.Product {
	// Simulacion
	fmt.Println("Obteniendo de MySQL")
	return m.products
}

func (m *MySQL) Update(product domain.Product) {
	for i, p := range m.products {
		if p.GetID() == product.GetID() {
			m.products[i] = product
			fmt.Println("Actualizando en MySQL")
			return
		}
	}
	fmt.Println("No se encontro el producto")
}

func (m *MySQL) Delete(id int32) {
	for i, p := range m.products {
		if p.GetID() == id {
			m.products = append(m.products[:i], m.products[i+1:]...)
			fmt.Println("Eliminando de MySQL")
			return
		}
	}
	fmt.Println("No se encontro el producto")
}
