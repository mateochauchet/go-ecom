package user

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/mateochauchet/go-ecom/config"
	"github.com/mateochauchet/go-ecom/services/auth"
	"github.com/mateochauchet/go-ecom/types"
	"github.com/mateochauchet/go-ecom/utils"
)

type Handler struct {
	//user repository
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{
		store,
	}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleRegister).Methods("POST")
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
	// get JSON payload
	var payload types.LoginUserPayload

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

	// check if user exists
	user, err := h.store.GetUserByEmail(payload.Email)

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("Not found, invalid email"))
		return
	}

	if !auth.ComparePassword(user.Password, []byte(payload.Password)) {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid email or password"))
		return
	}

	// generate JWT
	secret := []byte(config.Envs.JWTSecret)
	token, err := auth.CreateJWT([]byte(secret), user.ID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJson(w, http.StatusOK, map[string]string{"token": token})
}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	// get JSON payload
	var payload types.RegisterUserPayload

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

	// check if user already exists
	_, err = h.store.GetUserByEmail(payload.Email)

	if err == nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user already exists"))
		return
	}

	hashedPassword, err := auth.HashPassword(payload.Password)

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	// create user
	err = h.store.CreateUser(&types.User{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     payload.Email,
		Password:  hashedPassword,
	})

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJson(w, http.StatusCreated, map[string]string{"message": "user created"})
}
