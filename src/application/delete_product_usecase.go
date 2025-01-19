package application

import (
	"demo/src/domain"
)

type DeleteUseCase struct{
	db domain.IProduct
}

func NewUseCaseDelete(db domain.IProduct) *DeleteUseCase {
	return &DeleteUseCase{db:db}
}

func (uc *DeleteUseCase) Run(id int32) {
	uc.db.Delete(id)
}