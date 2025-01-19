package domain


type IProduct interface {
	Save(product Product)
	GetAll() []Product
	Update(product Product)
	Delete(id int32)
}