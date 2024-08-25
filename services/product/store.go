package product

import (
	"database/sql"
	"fmt"

	"github.com/mateochauchet/go-ecom/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) GetProductById(id int) (*types.Product, error) {
	return nil, nil
}

func (s *Store) GetProducts() ([]types.Product, error) {
	var query = "SELECT * FROM products"
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}

	products := make([]types.Product, 0)

	for rows.Next() {
		product, err := scanRowIntoProduct(rows)
		if err != nil {
			return nil, err
		}

		products = append(products, *product)
	}

	return products, nil
}

func (s *Store) CreateProduct(product *types.Product) error {
	var query = "INSERT INTO products (name, description, price, quantity, image) VALUES (?, ?, ?, ?, ?)"

	_, err := s.db.Exec(query, product.Name, product.Description, product.Price, product.Quantity, product.Image)

	if err != nil {
		return err
	}

	return nil
}

func scanRowIntoProduct(rows *sql.Rows) (*types.Product, error) {
	var product = new(types.Product)

	err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Quantity, &product.Image, &product.CreatedAt)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return product, nil
}
