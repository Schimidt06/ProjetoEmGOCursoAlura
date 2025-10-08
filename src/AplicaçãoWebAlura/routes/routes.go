// routes/routes.go

package routes

import (
	"aplicacaoWebAlura/controllers"
	"net/http"
)

func CarregaRotas() {
	http.HandleFunc("/", controllers.Index)
	// NOVAS ROTAS
	http.HandleFunc("/new", controllers.New)
	http.HandleFunc("/insert", controllers.Insert)
}