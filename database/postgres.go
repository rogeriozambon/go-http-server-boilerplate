package database

import (
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // isolated
)

var source = os.Getenv("DATABASE_URL")

type Postgres struct {
	DB *sqlx.DB
}

func NewPostgres() (p Postgres) {
	db, err := sqlx.Open("postgres", source)
	if err != nil {
		panic(err)
	}
	p.DB = db
	return
}
