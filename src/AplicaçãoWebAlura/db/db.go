package db

import (
	"database/sql"

	_ "github.com/lib/pq" // Driver do PostgreSQL
)

// A função começa com 'C' maiúsculo para ser "exportada" (pública).
// Assim, outros pacotes (como o 'main') podem enxergá-la.
func ConectaComBancoDeDados() *sql.DB {
	conexao := "user=postgres dbname=alura_loja password=061823 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}
	return db
}