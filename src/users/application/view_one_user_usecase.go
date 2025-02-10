package application

import "demo/src/users/domain"

type ViewOneUserUseCase struct {
	db domain.IUser
}

func NewViewOneUserUseCase(db domain.IUser) *ViewOneUserUseCase {
	return &ViewOneUserUseCase{db: db}
}

func (uc *ViewOneUserUseCase) Run(id int32) (domain.User, error) {
	return uc.db.GetByID(id)
}