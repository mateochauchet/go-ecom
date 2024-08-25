package product

import (
	"database/sql"

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
	return nil, nil
}
