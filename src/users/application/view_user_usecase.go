package application

import (
	"demo/src/users/domain"
)

type ViewUserUseCase struct {
	db domain.IUser
}

func NewUseCaseView(db domain.IUser) *ViewUserUseCase {
	return &ViewUserUseCase{db: db}
}

// Obtener a todos los usuarios
func (uc *ViewUserUseCase) Run() ([]domain.User, error) {
	return uc.db.GetAll()
}