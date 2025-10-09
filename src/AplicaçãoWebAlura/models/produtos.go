package models

import (
	"aplicacaoWebAlura/db"
	"log"      // IMPORT ADICIONADO
	"strconv"  // IMPORT ADICIONADO
)

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

// ... (todas as outras funções continuam iguais) ...
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

func DeletaProduto(id string) {
	db := db.ConectaComBancoDeDados()
	defer db.Close()

	deletaOProduto, err := db.Prepare("DELETE FROM produtos WHERE id=$1")
	if err != nil {
		panic(err.Error())
	}

	_, err = deletaOProduto.Exec(id)
	if err != nil {
		panic(err.Error())
	}
}

func BuscaUmProduto(id string) Produto {
	db := db.ConectaComBancoDeDados()
	defer db.Close()

	produtoDoBanco, err := db.Query("SELECT * FROM produtos WHERE id=$1", id)
	if err != nil {
		panic(err.Error())
	}

	produtoParaAtualizar := Produto{}

	for produtoDoBanco.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = produtoDoBanco.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}
		produtoParaAtualizar.Id = id
		produtoParaAtualizar.Nome = nome
		produtoParaAtualizar.Descricao = descricao
		produtoParaAtualizar.Preco = preco
		produtoParaAtualizar.Quantidade = quantidade
	}
	return produtoParaAtualizar
}


// AtualizaProduto atualiza um produto no banco de dados.
func AtualizaProduto(id string, nome string, descricao string, preco float64, quantidade int) {
	db := db.ConectaComBancoDeDados()
	defer db.Close()

	idConvertidoParaInt, err := strconv.Atoi(id)
	if err != nil {
		log.Println("Erro na conversão do ID para int:", err)
	}

	atualizaProduto, err := db.Prepare("UPDATE produtos SET nome=$1, descricao=$2, preco=$3, quantidade=$4 WHERE id=$5")
	if err != nil {
		panic(err.Error())
	}
	_, err = atualizaProduto.Exec(nome, descricao, preco, quantidade, idConvertidoParaInt)
	if err != nil {
		panic(err.Error())
	}
}