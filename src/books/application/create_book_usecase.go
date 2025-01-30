package application

import (
	"fmt"
	booksDomain "demo/src/books/domain"
)

type CreateBookUseCase struct {
	db booksDomain.IBook
}

func NewCreateUseCase(db booksDomain.IBook) *CreateBookUseCase {
	return &CreateBookUseCase{db: db}
}

func (uc *CreateBookUseCase) Run(book booksDomain.Book) (booksDomain.Book, error) {
	// Validar que el autor existe antes de guardar
	exists, err := uc.db.ValidateAuthor(book.Author)
	if err != nil {
		return booksDomain.Book{}, err
	}
	if !exists {
		return booksDomain.Book{}, fmt.Errorf("el autor con ID %d no existe", book.Author)
	}

	book, err = uc.db.Save(book)
	if err != nil {
		return booksDomain.Book{}, err
	}
	
	return book, nil
}
