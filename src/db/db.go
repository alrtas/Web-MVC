package db

import (
	"database/sql"

	//Fica como empty porque só é necessário no runtime
	_ "github.com/lib/pq"
)

func ContectaDb() *sql.DB {
	conexao := "user=postgres dbname=alura_loja password=admin host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao)

	if err != nil {
		panic(err.Error())
	}

	return db
}
