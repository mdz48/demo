package application

import (
	"demo/src/books/domain"
	"fmt"
)

type DeleteFavoriteUseCase struct {
	bookRepository domain.IBook
}

func NewDeleteFavoriteUseCase(bookRepository domain.IBook) *DeleteFavoriteUseCase {
	return &DeleteFavoriteUseCase{bookRepository: bookRepository}
}

func (u *DeleteFavoriteUseCase) Run(userId int32, bookId int32) error {
    rowsAffected, err := u.bookRepository.DeleteFavoriteBook(userId, bookId)
    if err != nil {
        return err
    }
    if rowsAffected == 0 {
        return fmt.Errorf("no se encontró el libro en favoritos")
    }
    return nil
}