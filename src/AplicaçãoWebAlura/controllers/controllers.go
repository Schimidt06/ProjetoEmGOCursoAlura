package controllers

import (
	"aplicacaoWebAlura/models"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

// Index carrega a lista de produtos
func Index(w http.ResponseWriter, r *http.Request) {
	produtos := models.BuscaTodosOsProdutos()
	temp.ExecuteTemplate(w, "Index", produtos)
}

// New renderiza o formulário de criação
func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

// Insert processa os dados do formulário e salva no banco
func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco, err := strconv.ParseFloat(r.FormValue("preco"), 64)
		if err != nil {
			log.Println("Erro na conversão do preço:", err)
		}

		quantidade, err := strconv.Atoi(r.FormValue("quantidade"))
		if err != nil {
			log.Println("Erro na conversão da quantidade:", err)
		}

		models.CriaNovoProduto(nome, descricao, preco, quantidade)
	}
	http.Redirect(w, r, "/", 301)
}