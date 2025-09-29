package main

import (
	"fmt"
)

func main() {
	nome := "João Schimidt"
	versao := 2.4
	fmt.Println("Olá, Sr.", nome)
	fmt.Println("Este programa está na versão", versao)

	fmt.Println("1-Inicar Monitoramento")
	fmt.Println("2-Exibir Logs")
	fmt.Println("0-Fechar essa bosta")

	var comando int
	fmt.Scan(&comando)
	fmt.Println(" O Endereço de memória da variavel comando é", &comando)
	fmt.Println("A variavel escolhida foi", comando)
}
