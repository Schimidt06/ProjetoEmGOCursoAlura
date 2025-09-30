package main

import (
	"fmt"
	"net/http"
	"os"
	"reflect"
)

func main() {
	var sites [4]string
	sites[0] = "https://random-status-code.herokuapp.com/"
	sites[1] = "https://www.alura.com.br"
	sites[2] = "https://www.caelum.com.br"
	fmt.Println(reflect.TypeOf(sites))
	fmt.Println(sites)
	exibeNomes()
	//exibeIntroducao()
	for {
		//exibeMenu()

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
	fmt.Println("Monitorando...")
	var sites [4]string
	sites[0] = "https://random-status-code.herokuapp.com/"
	sites[1] = "https://www.alura.com.br"
	sites[2] = "https://www.caelum.com.br"

	fmt.Println(sites)

	site := "https://random-status-code.herokuapp.com/"
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site:", site, "esta com problemas. Status Code:", resp.StatusCode)
	}
}

func exibeNomes() {
	nomes := []string{"João", "Maria", "José"}

	fmt.Println(nomes)
	fmt.Println(reflect.TypeOf(nomes))
	fmt.Println("O meu slice tem", len(nomes), "itens")
	fmt.Println("O meu slice tem capacidade para", cap(nomes), "itens")

	nomes = append(nomes, "Ana")

	fmt.Println(nomes)
	fmt.Println(reflect.TypeOf(nomes))
	fmt.Println("O meu slice tem", len(nomes))
	fmt.Println("O meu slice tem capacidade para", cap(nomes), "itens")
}
