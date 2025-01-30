package application

import (
	domain2 "demo/src/products/domain"
	"fmt"
)

// Como usare slice? Se tiene que usar struct

type ViewUseCase struct {
	db domain2.IProduct
}

func NewUseCaseCreate(db domain2.IProduct) *ViewUseCase {
	return &ViewUseCase{db: db}
}

func (uc *ViewUseCase) Run() []domain2.Product {
	products := uc.db.GetAll()
	for _, product := range products {
		fmt.Println(product)
	}
	return products
}
