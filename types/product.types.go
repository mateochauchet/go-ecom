package types

type Product struct {
}

type ProductStore interface {
	GetProductById(id int) (*Product, error)
	GetProducts() ([]Product, error)
}
