package domain

type FavoriteBookRequest struct {
    UserId int32 `json:"userId"`
    BookId int32 `json:"bookId"`
}