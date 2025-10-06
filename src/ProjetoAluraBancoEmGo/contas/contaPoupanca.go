// Define que este arquivo e tudo que está nele pertence ao pacote 'contas'.
// Agrupar tipos relacionados (como ContaCorrente e ContaPoupanca) no mesmo pacote
// é uma ótima prática de organização de código.
package contas

// Importamos o pacote 'clientes', que nós mesmos criamos.
// Isso é necessário para que possamos usar a struct 'Titular' que está definida lá.
// É um exemplo de como os pacotes se comunicam e reutilizam código em um projeto Go.
import "ProjetoAluraBancoEmGo/clientes"

// Aqui definimos a 'planta baixa' ou o 'molde' para criar uma Conta Poupança.
// Uma struct é um tipo que agrupa diferentes campos de dados em uma única unidade.
type ContaPoupanca struct {
	// CAMPO 1 (COMPOSIÇÃO): Em vez de copiar os campos (Nome, CPF...), nós 'embutimos'
	// a struct 'Titular' aqui. Isso significa que uma 'ContaPoupanca' TEM UM 'Titular'.
	// É a forma como Go favorece a composição em vez da herança.
	Titular clientes.Titular

	// CAMPOS 2, 3 e 4: Campos simples do tipo inteiro para guardar os dados da conta.
	NumeroAgencia, NumeroConta, Operacao int

	// CAMPO 5 (ENCAPSULAMENTO): O campo 'saldo' começa com letra minúscula.
	// Isso o torna PRIVADO e visível apenas para o código DENTRO do pacote 'contas'.
	// Nenhum código fora deste pacote (como o main.go) pode acessar 'c.saldo' diretamente.
	saldo float64
}

// Este é um MÉTODO da struct ContaPoupanca. Um método é uma função "atrelada" a um tipo específico.
// (c *ContaPoupanca) é o 'receiver' (receptor). O '*' indica que o método opera sobre um
// PONTEIRO para uma ContaPoupanca. Usamos um ponteiro porque este método
// PRECISA MODIFICAR o valor original da conta (o campo 'saldo').
func (c *ContaPoupanca) Sacar(valorDoSaque float64) (string, float64) {
	// Uma variável booleana que guarda a lógica de negócio para a operação de saque.
	// O saque só é válido se o valor for positivo E se houver saldo suficiente na conta.
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo

	// Verificamos o resultado da nossa lógica de negócio.
	if podeSacar {
		// MODIFICAÇÃO DO ESTADO: O valor do saque é subtraído do saldo da conta.
		// É por isso que o ponteiro no receiver é essencial. Sem ele, estaríamos
		// alterando uma cópia da conta, e não a original.
		c.saldo -= valorDoSaque
		// Retorna a mensagem de sucesso e o novo saldo atualizado.
		return "Saque realizado com sucesso", c.saldo
	} else {
		// Se a condição 'podeSacar' for falsa, a operação é negada.
		return "saldo insuficiente", c.saldo
	}
}

// Outro método para a ContaPoupanca, também com um ponteiro como receiver,
// pois ele também MODIFICA o saldo.
func (c *ContaPoupanca) Depositar(valorDoDeposito float64) (string, float64) {
	// Lógica de negócio: só permite o depósito de valores maiores que zero.
	if valorDoDeposito > 0 {
		// MODIFICAÇÃO DO ESTADO: O valor é adicionado ao saldo da conta.
		c.saldo += valorDoDeposito
		// Retorna a mensagem de sucesso e o novo saldo.
		return "Depósito realizado com sucesso", c.saldo
	} else {
		// Caso o valor seja inválido (zero ou negativo).
		return "Valor inválido para depósito", c.saldo
	}
}

// Este é um método "getter" ou "acessor". Sua única finalidade é expor de
// forma segura e controlada o valor do campo privado 'saldo'.
// Como 'saldo' é inacessível de fora do pacote, esta é a única maneira
// que o 'main.go' tem de consultar o saldo de uma conta.
func (c *ContaPoupanca) ObterSaldo() float64 {
	// Embora este método não modifique o saldo (apenas leia), é uma convenção
	// em Go manter todos os métodos de um tipo com o mesmo tipo de receiver
	// (ponteiro, neste caso) para maior consistência.
	return c.saldo
}