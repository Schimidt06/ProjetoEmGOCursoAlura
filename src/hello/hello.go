package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const monitoramentos = 5
const delay = 5 * time.Second

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
			os.Exit(255) // Use 255 para compatibilidade cross-plataforma
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
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
}

func leComando() int {
	var comandoLido int
	_, err := fmt.Scan(&comandoLido)
	if err != nil {
		return -1 // Garante que entradas inválidas caiam no default
	}
	fmt.Println("O comando escolhido foi", comandoLido)
	fmt.Println("")
	return comandoLido
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	sites := []string{
		"https://random-status-code.herokuapp.com/",
		"https://www.alura.com.br",
		"https://www.caelum.com.br",
	}

	for i := 0; i < monitoramentos; i++ {
		for j, site := range sites {
			fmt.Println("Testando site", j, ":", site)
			testaSite(site)
		}
		time.Sleep(delay) // Corrigido: delay já está em time.Duration
		fmt.Println("")
	}
	fmt.Println("")
}

func testaSite(site string) {
	resp, err := http.Get(site)
	if err != nil {
		fmt.Println("Erro ao acessar o site:", site)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site:", site, "esta com problemas. Status Code:", resp.StatusCode)
	}
}
