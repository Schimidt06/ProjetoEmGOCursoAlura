package main

import (
	"aplicacaoWebAlura/routes" // Importamos APENAS nosso pacote de rotas
	"log"
	"net/http"
)

func main() {
	// A única coisa que o main faz agora é carregar as rotas
	routes.CarregaRotas()

	// E iniciar o servidor.
	log.Println("Servidor rodando na porta 8000")
	http.ListenAndServe(":8000", nil)
}