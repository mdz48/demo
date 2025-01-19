package application

import (
	"demo/src/domain"
)

type UpdateUseCase struct{
	db domain.IProduct
}

func NewUseCaseUpdate(db domain.IProduct) *UpdateUseCase {
	return &UpdateUseCase{db:db}
}

func (uc *UpdateUseCase) Run(product domain.Product) {
	uc.db.Update(product)
}

