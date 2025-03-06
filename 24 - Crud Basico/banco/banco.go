package banco

import (
	"database/sql"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func Conectar() (*sql.DB, error) {
	urlConexao := "postgres://golang:golang@localhost:5432/devbook?sslmode=disable"
	db, erro := sql.Open("pgx", urlConexao)
	if erro != nil {
		return nil, erro
	}
	if erro = db.Ping(); erro != nil {
		db.Close()
		return nil, erro
	}
	return db, nil
}
