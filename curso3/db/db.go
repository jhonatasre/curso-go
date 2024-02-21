package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConectaBancoDados() *sql.DB {
	conexao := "user=postgres dbname=alura_loja password=admin host=localhost sslmode=disable port=15432"
	db, err := sql.Open("postgres", conexao)

	if err != nil {
		panic(err.Error())
	}

	return db
}
