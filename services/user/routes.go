package user

import (
	"fmt"
	"net/http"

	"github.com/Chandra5468/golangproject3/services/auth"
	"github.com/Chandra5468/golangproject3/types"
	"github.com/Chandra5468/golangproject3/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

// Each route service is going to be of type handler
// Where handler can take any dependecies
type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{
		store: store,
	}
}

// below method takes in the router
func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleRegister).Methods("POST")
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	var payload types.RegisterUserPayload

	if err := utils.ParseJson(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// validate the payload
	err := utils.Validate.Struct(payload)
	if err != nil {
		errors := err.(validator.ValidationErrors) // why do we use . after err.
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors))
		return
	}

	// check if user exists
	_, err = h.store.GetUserByEmail(payload.Email)
	if err == nil {
		utils.WriteError(w, http.StatusNotAcceptable, fmt.Errorf("user with email %s already exists", payload.Email))
		return
	}

	// hash or encode the password before saving it.
	hashedPassword, err := auth.HashPasswords(payload.Password)

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	// if it does not, create the new user
	err = h.store.CreateUser(types.User{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     payload.Email,
		Password:  hashedPassword,
	})

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJson(w, http.StatusCreated, nil)
}
