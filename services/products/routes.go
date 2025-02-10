package products

import (
	"net/http"

	"github.com/Chandra5468/golangproject3/types"
	"github.com/Chandra5468/golangproject3/utils"
	"github.com/gorilla/mux"
)

type Handler struct {
	store types.ProductStore
}

func NewHandler(store types.ProductStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/products", h.handlerCreateProduct).Methods("GET")
}

func (h *Handler) handlerCreateProduct(w http.ResponseWriter, r *http.Request) {
	ps, err := h.store.GetProducts()

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJson(w, http.StatusCreated, ps)
}
