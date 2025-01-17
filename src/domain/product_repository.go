package domain


type IProduct interface {
	Save(product Product)
	GetAll() []Product
}