package domain


type IProduct interface {
	Save(product Product) (Product, error)
	GetAll() []Product
	Update(product Product) (Product, error)
	Delete(id int32) (int64, error)
	GetByID(id int32) (Product, error)
}