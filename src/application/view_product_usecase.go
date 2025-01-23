package application

import (
	"demo/src/domain"
	"fmt"
)

// Como usare slice? Se tiene que usar struct

type ViewUseCase struct{
	db domain.IProduct
}

func NewUseCaseCreate(db domain.IProduct) *ViewUseCase {
	return &ViewUseCase{db:db}
}

func (uc *ViewUseCase) Run() {
	products := uc.db.GetAll()
	for _, product := range products {
		fmt.Println(product)
	}
}