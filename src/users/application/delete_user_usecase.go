package application

import (
	"demo/src/users/domain"
)

type DeleteUseCase struct {
	db domain.IUser
}

func NewUseCaseDelete(db domain.IUser) *DeleteUseCase {
	return &DeleteUseCase{db: db}
}

func (uc *DeleteUseCase) Run(id int32) (int64, error) {
	result, err := uc.db.Delete(id)
	return result, err
}