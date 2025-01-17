package application

import "demo/src/domain"

type CreateProductUseCase struct{
	db domain.IProduct
}

func NewCreateUseCase(db domain.IProduct) *CreateProductUseCase {
	return &CreateProductUseCase{db: db}
}

func (uc *CreateProductUseCase) Run(product domain.Product) {
	uc.db.Save(product)
}

