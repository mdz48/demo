package domain

type FavoriteBookRequest struct {
    UserId int32 `json:"user_id"`
    BookId int32 `json:"book_id"`
}