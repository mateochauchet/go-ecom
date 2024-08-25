package product

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mateochauchet/go-ecom/types"
)

type handler struct {
	store types.ProductStore
}

func NewHandler(store types.ProductStore) *handler {
	return &handler{
		store,
	}
}

func (h *handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/products", h.handleGetProducts).Methods("GET")
}

func (h *handler) handleGetProducts(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetProducts")
}
