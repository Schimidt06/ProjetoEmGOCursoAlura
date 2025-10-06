package main

import (
	"ProjetoAluraBancoEmGo/contas"
	"fmt"
)

// A INTERFACE FOI CORRIGIDA AQUI: agora espera o retorno de (string, float64)
type verificarConta interface {
	Sacar(valor float64) (string, float64)
}

// A FUNÇÃO FOI CORRIGIDA AQUI: agora ela lida com os retornos do Sacar
func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	status, _ := conta.Sacar(valorDoBoleto) // Usamos o blank identifier _ para ignorar o saldo retornado
	fmt.Println("Status do Pagamento do Boleto:", status)
}

func main() {
	contaDoDenis := contas.ContaPoupanca{}
	contaDoDenis.Depositar(100)
	PagarBoleto(&contaDoDenis, 60)

	fmt.Println("Saldo final:", contaDoDenis.ObterSaldo())

	contaDaLaura := contas.ContaCorrente{}
	contaDaLaura.Depositar(200)
	PagarBoleto(&contaDaLaura, 80)
	fmt.Println("Saldo final:", contaDaLaura.ObterSaldo())
}
