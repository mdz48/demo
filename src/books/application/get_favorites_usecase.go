package application

import "demo/src/books/domain"

type GetFavoritesUseCase struct {
	bookRepository domain.IBook
}

func NewGetFavoritesUseCase(bookRepository domain.IBook) *GetFavoritesUseCase {
	return &GetFavoritesUseCase{bookRepository: bookRepository}
}

func (u *GetFavoritesUseCase) Run(userId int32) ([]domain.BookWithAuthor, error) {
	return u.bookRepository.GetFavoriteBooks(userId)
}
