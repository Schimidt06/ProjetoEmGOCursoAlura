package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
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
			imprimeLogs()
		case 0:
			fmt.Println("Fechando o programa...")
			os.Exit(0)
		default:
			fmt.Println("Não conheço este comando")
			// Decidi não sair mais com os.Exit para o usuário poder tentar de novo
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
	fmt.Println("\n1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
}

func leComando() int {
	reader := bufio.NewReader(os.Stdin)
	// Lê a string até o caractere de nova linha '\n'
	input, _ := reader.ReadString('\n')
	// Remove espaços em branco e a quebra de linha
	comandoStr := strings.TrimSpace(input)
	// Converte a string limpa para inteiro
	comando, err := strconv.Atoi(comandoStr)

	// Se der erro na conversão (ex: usuário digitou "abc"), retorna um comando inválido
	if err != nil {
		fmt.Println("Por favor, digite um número válido.")
		return -1 // Um número que não está no seu menu para cair no default
	}

	fmt.Println("O comando escolhido foi", comando)
	fmt.Println("")
	return comando
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	sites := leSitesDoArquivo()

	if len(sites) == 0 {
		fmt.Println("Nenhum site encontrado no arquivo 'sites.txt'.")
		return
	}

	for i := 0; i < monitoramentos; i++ {
		fmt.Println("--- Ciclo de Monitoramento", i+1, "---")
		for j, site := range sites {
			fmt.Println("Testando site", j, ":", site)
			testaSite(site)
		}
		time.Sleep(delay)
		fmt.Println("")
	}
	fmt.Println("")
}

func testaSite(site string) {
	resp, err := http.Get(site)
	if err != nil {
		fmt.Println("Erro ao acessar o site:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
		registraLog(site, true)
	} else {
		fmt.Println("Site:", site, "esta com problemas. Status Code:", resp.StatusCode)
		registraLog(site, false)
	}
}

func leSitesDoArquivo() []string {
	var sites []string

	arquivo, err := os.Open("sites.txt")
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo 'sites.txt':", err)
		return sites
	}
	defer arquivo.Close()

	scanner := bufio.NewScanner(arquivo)
	for scanner.Scan() {
		linha := strings.TrimSpace(scanner.Text())
		if linha != "" {
			sites = append(sites, linha)
		}
	}
	return sites
}

func registraLog(site string, status bool) {
	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}
	defer arquivo.Close()

	statusStr := "online: false"
	if status {
		statusStr = "online: true"
	}

	logLine := fmt.Sprintf("%s - %s - %s\n", time.Now().Format("02/01/2006 15:04:05"), site, statusStr)
	arquivo.WriteString(logLine)
}

func imprimeLogs() {
	arquivo, err := os.ReadFile("log.txt")
	if err != nil {
		fmt.Println("Erro ao ler o arquivo de log:", err)
		return
	}
	fmt.Println("--- INÍCIO DOS LOGS ---")
	fmt.Print(string(arquivo))
	fmt.Println("--- FIM DOS LOGS ---")
}
