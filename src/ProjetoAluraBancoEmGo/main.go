package main

import (
	"ProjetoAluraBancoEmGo/contas"
	"fmt"
)

func main() {
	contaExemplo := contas.ContaCorrente{}
	contaExemplo.Depositar(100)

	fmt.Println(contaExemplo.ObterSaldo())
}
