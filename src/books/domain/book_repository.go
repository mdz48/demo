package domain

type IBook interface {
	Save(book Book) (Book, error)
	ValidateAuthor(authorId int32) (bool, error)
	Update(id int32 , book Book) (Book, error)
	Delete(id int32) (int64, error)
	GetAll() ([]BookWithAuthor, error)
	GetBooksByAuthor(authorId int32) ([]BookWithAuthor, error)
}
