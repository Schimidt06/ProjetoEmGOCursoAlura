package main

import "fmt"
import "os"
	
	

func main() {

	exibeIntroducao()
	exibeMenu()
	comando := leComando()

	switch comando {
	case 1:
		fmt.Println("Monitoramento iniciado...")
	case 2:
		fmt.Println("Exibindo Logs...")
	case 0:
		fmt.Println("Fechando o programa...")
		os.Exit(0)
	default:
		fmt.Println("Comando inválido")
		os.Exit(-1)
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
	fmt.Println("A variavel escolhida foi", comandoLido)
	return comandoLido
}
