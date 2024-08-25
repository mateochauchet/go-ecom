package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mateochauchet/go-ecom/services/product"
	"github.com/mateochauchet/go-ecom/services/user"
)

type APIServer struct {
	address string
	db      *sql.DB
}

func NewAPIServer(address string, db *sql.DB) *APIServer {
	return &APIServer{
		address: address,
		db:      db,
	}
}

func (s *APIServer) Run() error {
	var router = mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	router.HandleFunc("/ping", handlePing).Methods("GET")

	// Register user routes
	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(subrouter)

	// Register product routes
	productStore := product.NewStore(s.db)
	productHandler := product.NewHandler(productStore)
	productHandler.RegisterRoutes(subrouter)

	log.Println("Listening on port " + s.address)

	return http.ListenAndServe(s.address, router)
}

func handlePing(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}
