package application

import "demo/src/users/domain"

type LoginUserUseCase struct {
	db domain.IUser
}

func NewUseCaseLogin(db domain.IUser) *LoginUserUseCase {
	return &LoginUserUseCase{db: db}
}

func (u *LoginUserUseCase) Login(email string, password string) (bool, error) {
	return u.db.Login(email, password)
}
