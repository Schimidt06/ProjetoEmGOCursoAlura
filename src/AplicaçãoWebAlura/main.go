package main

import (
	"html/template"
	"log"
	"net/http"

	"aplicacaoWebAlura/db"
	"aplicacaoWebAlura/models" // <-- NOVA IMPORTAÇÃO
)

// A struct Produto FOI REMOVIDA DAQUI e movida para models/produtos.go

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	log.Println("Servidor rodando na porta 8000")
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	db := db.ConectaComBancoDeDados()
	defer db.Close()

	selectDeTodosOsProdutos, err := db.Query("SELECT * FROM produtos ORDER BY id ASC")
	if err != nil {
		panic(err.Error())
	}

	// Agora usamos "models.Produto" para referenciar a struct
	produtos := []models.Produto{}

	for selectDeTodosOsProdutos.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectDeTodosOsProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		// Ao criar o produto, também usamos o prefixo do pacote "models"
		p := models.Produto{
			Id:         id,
			Nome:       nome,
			Descricao:  descricao,
			Preco:      preco,
			Quantidade: quantidade,
		}

		produtos = append(produtos, p)
	}

	temp.ExecuteTemplate(w, "Index", produtos)
}