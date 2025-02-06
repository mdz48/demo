package application

import "demo/src/books/domain"

type AddFavoriteBookUseCase struct {
    db domain.IBook
}

func NewAddFavoriteBookUseCase(db domain.IBook) *AddFavoriteBookUseCase {
    return &AddFavoriteBookUseCase{db: db}
}

func (uc *AddFavoriteBookUseCase) Run(userId int32, bookId int32) error {
    _, err := uc.db.AddFavoriteBook(userId, bookId)
    return err
}