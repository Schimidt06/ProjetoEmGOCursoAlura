package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {
	//primeiro modo de fazer a varavel
	//contaDoJoao := ContaCorrente{titular: "João", numeroAgencia: 123, numeroConta: 456789, saldo: 1000.0}
	//segundo modo de fazer a variavel

	//contaDaBruna2:= ContaCorrente{"Bruna", 321, 987654, 2500.0}

	//fmt.Println(contaDoJoao == contaDoJoao2)
	//fmt.Println(contaDaBruna)

	//Terceiro modo de fazer a variavel
	var contaDaCristina *ContaCorrente
	contaDaCristina = new(ContaCorrente)
	contaDaCristina.titular = "Cristina"
	contaDaCristina.saldo = 3000.0

	var contaDaCristina2 *ContaCorrente
	contaDaCristina2 = new(ContaCorrente)
	contaDaCristina2.titular = "Cristina"
	contaDaCristina2.saldo = 3000.0

	fmt.Println(&contaDaCristina)
	fmt.Println(&contaDaCristina2)

	fmt.Println(*contaDaCristina == *contaDaCristina2)
}
