package application

import (
	"demo/src/domain"
	"fmt"
)

type UpdateUseCase struct{
	db domain.IProduct
}

func NewUseCaseUpdate(db domain.IProduct) *UpdateUseCase {
	return &UpdateUseCase{db:db}
}

func (uc *UpdateUseCase) Run(id int32, product domain.Product) (domain.Product, error) {
	result, err := uc.db.GetByID(id)
	if err != nil {
		return domain.Product{}, err
	}
	if result.Id == 0 {
		return domain.Product{}, fmt.Errorf("product not found")
	}
	product.Id = id 
	result, err = uc.db.Update(product)
	if err != nil {
		return domain.Product{}, err
	}
	return result, err
}

