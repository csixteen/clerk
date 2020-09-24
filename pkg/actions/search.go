package actions

import (
	"database/sql"
)

type Result interface {
	String() string
}

func Search(db *sql.DB, args []string) ([]Result, error) {
	return nil, nil
}
