package cart

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mateochauchet/go-ecom/types"
	"github.com/mateochauchet/go-ecom/utils"
)

type handler struct {
	store types.CartStore
}

func NewHandler(store types.CartStore) *handler {
	return &handler{
		store,
	}
}

func (h *handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/cart/checkout", h.handleCheckout).Methods("POST")
}

func (h *handler) handleCheckout(w http.ResponseWriter, r *http.Request) {
	utils.WriteJson(w, http.StatusOK, map[string]string{"message": "Checkout"})
}
