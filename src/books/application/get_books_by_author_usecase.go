package application

import "demo/src/books/domain"

type ViewBooksByAuthorUseCase struct {
	db domain.IBook
}

func NewViewBooksByAuthorUseCase(db domain.IBook) *ViewBooksByAuthorUseCase {
	return &ViewBooksByAuthorUseCase{db: db}
}

func (uc *ViewBooksByAuthorUseCase) Run(authorId int32) ([]domain.Book, error) {
	return uc.db.GetBooksByAuthor(authorId)
}
