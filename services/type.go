package services

import (
	"encoding/json"
	"net/http"

	"github.com/jmoiron/sqlx"
)

// Services struct
type Services struct {
	db *sqlx.DB
}

// New initializes service struct
func New(db *sqlx.DB) Services {
	return Services{db}
}

// Register to setting up routes
func (s Services) Register(mux *http.ServeMux) {
	mux.HandleFunc("/", s.Home)
}

func (s Services) render(w http.ResponseWriter, statusCode int, response interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}
