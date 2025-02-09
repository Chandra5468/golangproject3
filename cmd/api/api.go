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

	// initializing a new router using mux
	router := mux.NewRouter()
	// adding prefix to all the routes. And return a subrouter. which will match the prefix pattern of all api requests.
	subrouter := router.PathPrefix("/v1/api/").Subrouter()

	// Below NewStore is method of storage file. Where all sql related ops happen
	userStore := user.NewStore(s.db)
	// userStore is a struct of sql.db fields. AND which has all db ops methods we created
	// we are passing the struct userStore(which has sql.db) and all db ops to NewHandler, which accepts interface.
	userServiceHandler := user.NewHandler(userStore) // IMP Implementing interface for userStore struct
	// This "NewHandler" return interface with method signatures of db ops of userStore
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
