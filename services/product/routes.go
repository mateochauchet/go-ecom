package product

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
	"github.com/mateochauchet/go-ecom/types"
	"github.com/mateochauchet/go-ecom/utils"
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
	router.HandleFunc("/products", h.CreateProduct).Methods("POST")
}

func (h *handler) handleGetProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.store.GetProducts()

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJson(w, http.StatusOK, products)
}

func (h *handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	// get JSON payload
	var payload types.CreateProductPayload

	err := utils.ParseJson(r, &payload)

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// validate payload
	err = utils.Validate.Struct(payload)

	if err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	// create user
	err = h.store.CreateProduct(&types.Product{
		Name:        payload.Name,
		Description: payload.Description,
		Price:       payload.Price,
		Image:       payload.Image,
		Quantity:    payload.Quantity,
	})

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJson(w, http.StatusCreated, map[string]string{"message": "product created"})
}
