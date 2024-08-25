package types

type Product struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Quantity    int    `json:"quantity"`
	Image       string `json:"image"`
	CreatedAt   string `json:"createdAt"`
}

type ProductStore interface {
	GetProductById(id int) (*Product, error)
	GetProducts() ([]Product, error)
	CreateProduct(product *Product) error
}

type CreateProductPayload struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	Price       int    `json:"price" validate:"required"`
	Quantity    int    `json:"quantity" validate:"required"`
	Image       string `json:"image" validate:"required"`
}
