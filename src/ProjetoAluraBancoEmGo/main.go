// Define que este arquivo pertence ao pacote 'main'. Em Go, o pacote 'main' é especial:
// ele indica que este código, quando compilado, vai gerar um programa executável.
// Pense nisso como a "ignição" do seu aplicativo.
package main

// 'import' é como trazemos "caixas de ferramentas" de outros pacotes para usar neste arquivo.
import (
	// Estamos importando nosso próprio pacote 'contas', que criamos dentro do projeto.
	// Isso nos dá acesso a tudo que for PÚBLICO (com letra maiúscula) dentro dele,
	// como as structs 'ContaPoupanca' e 'ContaCorrente' e seus métodos.
	"ProjetoAluraBancoEmGo/contas"

	// 'fmt' é um pacote que já vem com o Go. Ele contém funções para formatação
	// e impressão de texto, sendo a 'Println' (Print Line) a mais famosa.
	"fmt"
)

// Aqui definimos um 'contrato' ou um 'molde' de comportamento chamado 'verificarConta'.
// Uma interface não guarda dados, ela apenas define quais MÉTODOS um tipo precisa ter.
// É a base para criar sistemas flexíveis onde as partes não precisam se conhecer
// diretamente, apenas precisam respeitar o mesmo 'contrato'.
type verificarConta interface {
	// Este contrato exige uma única coisa: qualquer tipo que queira ser uma 'verificarConta'
	// precisa, obrigatoriamente, ter um método chamado 'Sacar'.
	// A assinatura deve ser EXATA: ele precisa receber um float64 e retornar uma (string, float64).
	Sacar(valor float64) (string, float64)
}

// Esta é uma função POLIMÓRFICA. A "mágica" das interfaces acontece aqui.
// O primeiro parâmetro 'conta' não é de um tipo específico (como ContaPoupanca),
// mas sim da interface 'verificarConta'. Isso significa que esta função pode receber
// QUALQUER tipo de variável que satisfaça aquele contrato que definimos acima.
func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	// Chamamos o método 'Sacar' do tipo que foi passado. Não importa se é uma
	// ContaPoupanca ou ContaCorrente, Go garante que o método 'Sacar' existe,
	// pois essa é a regra da interface.
	// O método Sacar retorna dois valores. Capturamos o primeiro (a string de status)
	// na variável 'status'. O segundo valor (o novo saldo) é deliberadamente ignorado
	// usando o 'blank identifier' ou 'identificador em branco' (_), pois não precisamos dele aqui.
	status, _ := conta.Sacar(valorDoBoleto)

	// Usamos o pacote 'fmt' para imprimir o status da operação no console.
	fmt.Println("Status do Pagamento do Boleto:", status)
}

// A função 'main' é o coração do programa, onde a execução realmente começa.
func main() {
	// --- Trabalhando com a Conta Poupança ---

	// Criamos uma instância (uma 'variável') do tipo 'ContaPoupanca'.
	// 'contaDoDenis' agora é uma variável que contém todos os campos de uma ContaPoupanca.
	contaDoDenis := contas.ContaPoupanca{}

	// Chamamos o método 'Depositar' da contaDoDenis para adicionar 100 ao seu saldo.
	contaDoDenis.Depositar(100)

	// Chamamos nossa função polimórfica, passando o ENDEREÇO DE MEMÓRIA (&) da contaDoDenis.
	// Go verifica se o tipo '*contas.ContaPoupanca' (um ponteiro para ContaPoupanca)
	// satisfaz a interface 'verificarConta'. Como ele tem o método Sacar com a assinatura
	// exata, a verificação passa e a função é executada para a conta do Denis.
	PagarBoleto(&contaDoDenis, 60)

	// Imprime o saldo final da contaDoDenis para vermos o resultado (deve ser 40).
	fmt.Println("Saldo final:", contaDoDenis.ObterSaldo())

	// --- Demonstração do Polimorfismo com a Conta Corrente ---

	// Agora, criamos uma instância de um tipo COMPLETAMENTE DIFERENTE: 'ContaCorrente'.
	contaDaLaura := contas.ContaCorrente{}

	// Depositamos 200 na conta da Laura.
	contaDaLaura.Depositar(200)

	// USAMOS EXATAMENTE A MESMA FUNÇÃO 'PagarBoleto'!
	// Não precisamos escrever uma 'PagarBoletoParaContaCorrente'. A função original é
	// reutilizável graças à interface. Go faz a mesma verificação e vê que
	// '*contas.ContaCorrente' também satisfaz o contrato.
	PagarBoleto(&contaDaLaura, 80)

	// Imprimimos o saldo final da contaDaLaura para confirmar o resultado (deve ser 120).
	fmt.Println("Saldo final:", contaDaLaura.ObterSaldo())
}