package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/mateochauchet/go-ecom/types"
)

func TestUserServiceHandler(t *testing.T) {
	var mockUserStore = &mockUserStore{}
	var handler = NewHandler(mockUserStore)

	t.Run("should fail is the user payload is invalid", func(t *testing.T) {
		var payload = types.RegisterUserPayload{
			FirstName: "Leonardo",
			LastName:  "da Vinci",
			Email:     "invalid email",
			Password:  "12345",
		}
		body, _ := json.Marshal(payload)

		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(body))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/register", handler.handleRegister)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("handler returned wrong status code: got %v want %v", rr.Code, http.StatusBadRequest)
		}
	})

	t.Run("should correctly register a new user", func(t *testing.T) {
		var payload = types.RegisterUserPayload{
			FirstName: "Leonardo",
			LastName:  "da Vinci",
			Email:     "3OqgV@example.com",
			Password:  "12345",
		}
		body, _ := json.Marshal(payload)

		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(body))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/register", handler.handleRegister)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusCreated {
			t.Errorf("handler returned wrong status code: got %v want %v", rr.Code, http.StatusCreated)
		}
	})

}

type mockUserStore struct{}

func (m *mockUserStore) GetUserByEmail(email string) (*types.User, error) {
	return nil, fmt.Errorf("user not found")
}

func (m *mockUserStore) GetUserById(id int) (*types.User, error) {
	return nil, nil
}

func (m *mockUserStore) CreateUser(user *types.User) error {
	return nil
}
