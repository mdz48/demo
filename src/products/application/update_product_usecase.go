package application

import (
	domain2 "demo/src/products/domain"
	"fmt"
)

type UpdateUseCase struct {
	db domain2.IProduct
}

func NewUseCaseUpdate(db domain2.IProduct) *UpdateUseCase {
	return &UpdateUseCase{db: db}
}

func (uc *UpdateUseCase) Run(id int32, product domain2.Product) (domain2.Product, error) {
	result, err := uc.db.GetByID(id)
	if err != nil {
		return domain2.Product{}, err
	}
	if result.Id == 0 {
		return domain2.Product{}, fmt.Errorf("product not found")
	}
	product.Id = id
	result, err = uc.db.Update(product)
	if err != nil {
		return domain2.Product{}, err
	}
	return result, err
}
