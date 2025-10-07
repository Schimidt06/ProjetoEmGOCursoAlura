package models

/*
   ATENÇÃO A UM DETALHE CRUCIAL DO GO!

   Para que os campos da struct (Id, Nome, Preco, etc.) possam ser acessados
   de fora do pacote 'models' (ou seja, pelo nosso 'main.go' e pelo sistema
   de templates), seus nomes DEVEM começar com letra maiúscula.
   É assim que o Go define o que é "público" (exportado) e "privado".
*/

// Produto representa a estrutura de um produto da loja.
type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}