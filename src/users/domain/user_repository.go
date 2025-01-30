package domain

type IUser interface {
	Save(user User) (User, error)
	GetAll() ([]User, error)
	Update(id int32 , user User) (User, error)
	Delete(id int32) (int64, error)
	GetByID(id int32) (User, error)
}