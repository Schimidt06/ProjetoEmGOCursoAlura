package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	exibeIntroducao()
	for {

		exibeMenu()

		comando := leComando()

		switch comando {
		case 1:

			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo Logs...")
		case 0:
			fmt.Println("Fechando o programa...")
			os.Exit(0)
		default:
			fmt.Println("Não conheço este comando")
			os.Exit(-1)
		}

	}

}

func exibeIntroducao() {
	nome := "João Schimidt"
	versao := 2.4
	fmt.Println("Olá, Sr.", nome)
	fmt.Println("Este programa está na versão", versao)
}
func exibeMenu() {
	fmt.Println("1-Inicar Monitoramento")
	fmt.Println("2-Exibir Logs")
	fmt.Println("0-Fechar essa bosta")
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	// fmt.Println("A variavel escolhida foi", comandoLido) // Remova ou comente esta linha
	return comandoLido
}

func iniciarMonitoramento() {
	fmt.Println("Monitoramento iniciado...")
	site := "https://random-status-code.herokuapp.com/"
	resp, _ := http.Get(site)
	// fmt.Println(resp) // Remova ou comente esta linha

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
	}
}
