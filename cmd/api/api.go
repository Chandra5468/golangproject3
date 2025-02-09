package api

import (
	"database/sql"
	"log/slog"
	"net/http"

	"github.com/Chandra5468/golangproject3/services/user"
	"github.com/gorilla/mux"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewApiServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	// router := http.NewServeMux()

	// router.HandleFunc("GET /v1/api/", DemoHandler)

	// Above router is golang inbuilt router.
	// using gorilla mux from below

	router := mux.NewRouter()
	subrouter := router.PathPrefix("/v1/").Subrouter()

	userStore := user.NewStore(s.db)
	userServiceHandler := user.NewHandler(userStore)

	userServiceHandler.RegisterRoutes(subrouter)
	slog.Info("message", "Listening on address", s.addr)
	return http.ListenAndServe(s.addr, router)
}

// func DemoHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Add("Content-Type", "application/json")
// 	w.Header().Add("Cache-Control", "max-age=0")
// 	w.WriteHeader(200)
// 	json.NewEncoder(w).Encode("This is response from backend")
// }
