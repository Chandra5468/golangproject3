package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Chandra5468/golangproject3/types"
	"github.com/gorilla/mux"
)

func TestUserServiceHandlers(t *testing.T) {
	userStore := &mockUserStore{}
	handler := NewHandler(userStore)

	t.Run("Should fail if user payload is invalid", func(t *testing.T) {
		payload := types.RegisterUserPayload{
			FirstName: "user",
			LastName:  "123",
			Email:     "xyz", // this is an invalid pattern
			Password:  "12",
		}

		marshelled, _ := json.Marshal(payload) // what does marshal do ?
		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshelled))

		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder() // what does httptest do ? and what is new recorder
		router := mux.NewRouter()

		router.HandleFunc("/register", handler.handleRegister)

		router.ServeHTTP(rr, req) //

		if rr.Code != http.StatusBadRequest { // what does Code do ?
			t.Errorf("expected status code %d, got %d", http.StatusBadRequest, rr.Code)
		}
	})

	t.Run("should correctly register the user", func(t *testing.T) {
		payload := types.RegisterUserPayload{
			FirstName: "user",
			LastName:  "123",
			Email:     "valid@gmail.com",
			Password:  "12",
		}

		marshelled, _ := json.Marshal(payload) // what does marshal do ?
		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshelled))

		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder() // what does httptest do ? and what is new recorder
		router := mux.NewRouter()

		router.HandleFunc("/register", handler.handleRegister)

		router.ServeHTTP(rr, req) //

		if rr.Code != http.StatusCreated { // what does Code do ?
			t.Errorf("expected status code %d, got %d", http.StatusCreated, rr.Code)
		}
	})
}

// we are replicating below struct and methods as types file(types.go) to test.
type mockUserStore struct {
}

func (m *mockUserStore) GetUserByEmail(email string) (*types.User, error) {
	return nil, fmt.Errorf("user not found")
}

func (m *mockUserStore) GetUserByID(id int) (*types.User, error) {
	return nil, nil
}

func (m *mockUserStore) CreateUser(types.User) error {
	return nil
}
