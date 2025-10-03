package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	} else {
		return "Saldo insuficiente"
	}
}

func main() {
	contaDaSilvia := ContaCorrente{}
	contaDaSilvia.titular = "Silvia"
	contaDaSilvia.saldo = 1000

	// Imprime o saldo ANTES do saque
	fmt.Println(contaDaSilvia.saldo)

	// Imprime o saldo DEPOIS do saque
	fmt.Println(contaDaSilvia.Sacar(500))
	fmt.Println(contaDaSilvia.saldo)
}
