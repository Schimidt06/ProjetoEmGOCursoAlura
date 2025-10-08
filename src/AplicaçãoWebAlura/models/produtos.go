package models

import (
	"aplicacaoWebAlura/db"
)

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

// BuscaTodosOsProdutos busca todos os produtos no banco de dados.
func BuscaTodosOsProdutos() []Produto {
	db := db.ConectaComBancoDeDados()
	defer db.Close()

	selectDeTodosOsProdutos, err := db.Query("SELECT * FROM produtos ORDER BY id ASC")
	if err != nil {
		panic(err.Error())
	}

	produtos := []Produto{}
	for selectDeTodosOsProdutos.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectDeTodosOsProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}
		p := Produto{
			Id:         id,
			Nome:       nome,
			Descricao:  descricao,
			Preco:      preco,
			Quantidade: quantidade,
		}
		produtos = append(produtos, p)
	}
	return produtos
}

// CriaNovoProduto insere um produto no banco
func CriaNovoProduto(nome, descricao string, preco float64, quantidade int) {
	db := db.ConectaComBancoDeDados()
	defer db.Close()

	insereDadosNoBanco, err := db.Prepare("INSERT INTO produtos (nome, descricao, preco, quantidade) VALUES ($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}
	_, err = insereDadosNoBanco.Exec(nome, descricao, preco, quantidade)
	if err != nil {
		panic(err.Error())
	}
}
