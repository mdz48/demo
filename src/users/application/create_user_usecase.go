package application

import (
	domain2 "demo/src/users/domain"
)

type CreateUserUseCase struct {
	db domain2.IUser
}

func NewCreateUseCase(db domain2.IUser) *CreateUserUseCase {
	return &CreateUserUseCase{db: db}
}

func (uc *CreateUserUseCase) Run(user domain2.User) domain2.User {
	user, _ = uc.db.Save(user)
	return user
}