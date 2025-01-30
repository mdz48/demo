package application

import (
	"demo/src/books/domain"
)

type ViewBooksUseCase struct {
	bookRepository domain.IBook
}

func NewUseCaseViewBooks(bookRepository domain.IBook) *ViewBooksUseCase {
	return &ViewBooksUseCase{bookRepository: bookRepository}
}

func (uc *ViewBooksUseCase) Run() ([]domain.Book, error) {
	return uc.bookRepository.GetAll()
}