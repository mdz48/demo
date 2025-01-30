package application

import (
	"demo/src/books/domain"
)

type UpdateUseCase struct {
	bookRepository domain.IBook
}

func NewUseCaseUpdate(bookRepository domain.IBook) *UpdateUseCase {
	return &UpdateUseCase{bookRepository: bookRepository}
}

func (uc *UpdateUseCase) Run(id int32, book domain.Book) (domain.Book, error) {
	return uc.bookRepository.Update(id, book)
}