package main

import (
	"ProjetoAluraBancoEmGo/contas"
	"fmt"
)

func main() {
	// Correção 1: Adicionado o prefixo "contas."
	// Correção 2: Campos com letra maiúscula (Titular, Saldo)
	contaDaSilvia := contas.ContaCorrente{Titular: "Silvia", Saldo: 300}
	contaDoGustavo := contas.ContaCorrente{Titular: "Gustavo", Saldo: 500}

	// Lembre-se de usar um valor válido para a transferência!
	// Ex: um valor positivo e que Gustavo tenha em conta.
	status := contaDoGustavo.Transferir(150, &contaDaSilvia)

	fmt.Println(status)
	fmt.Println(contaDaSilvia)
	fmt.Println(contaDoGustavo)
}
