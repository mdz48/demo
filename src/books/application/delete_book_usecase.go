package application

import (
	"demo/src/books/domain"
)

type DeleteUseCase struct {
	db domain.IBook
}

func NewUseCaseDelete(db domain.IBook) *DeleteUseCase {
	return &DeleteUseCase{db: db}
}

func (uc *DeleteUseCase) Run(id int32) (int64, error) {
	result, err := uc.db.Delete(id)
	return result, err
}