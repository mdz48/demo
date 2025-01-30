package application

import (
    "demo/src/users/domain"
    "fmt"
)

type UpdateUseCase struct {
    db domain.IUser
}

func NewUseCaseUpdate(db domain.IUser) *UpdateUseCase {
    return &UpdateUseCase{db: db}
}

func (uc *UpdateUseCase) Run(id int32, user domain.User) (domain.User, error) {
    // Verificar si existe el usuario
    existingUser, err := uc.db.GetByID(id)
    if err != nil {
        return domain.User{}, err
    }
    if existingUser.Id == 0 {
        return domain.User{}, fmt.Errorf("usuario no encontrado")
    }

    // Mantener el ID original
    user.Id = id

    // Actualizar usuario
    result, err := uc.db.Update(id, user)
    if err != nil {
        return domain.User{}, err
    }

    return result, nil
}