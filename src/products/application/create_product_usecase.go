package application

import (
	domain2 "demo/src/products/domain"
)

type CreateProductUseCase struct {
	db domain2.IProduct
}

func NewCreateUseCase(db domain2.IProduct) *CreateProductUseCase {
	return &CreateProductUseCase{db: db}
}

func (uc *CreateProductUseCase) Run(product domain2.Product) domain2.Product {
	product, _ = uc.db.Save(product)
	return product
}
