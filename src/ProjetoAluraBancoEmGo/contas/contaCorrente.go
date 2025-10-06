// Novamente, este arquivo faz parte do pacote 'contas', mantendo a organização do projeto.
package contas

// Importamos o pacote 'clientes' para usar a struct 'Titular'.
import "ProjetoAluraBancoEmGo/clientes"

// Definição da struct 'ContaCorrente'. Note que ela é ligeiramente diferente
// da ContaPoupanca (não possui o campo 'Operacao', por exemplo).
// Tipos diferentes podem ter estruturas de dados diferentes.
type ContaCorrente struct {
	// COMPOSIÇÃO: A ContaCorrente também TEM UM Titular.
	Titular clientes.Titular

	// Campos para os dados da agência e número da conta.
	NumeroAgencia, NumeroConta int

	// ENCAPSULAMENTO: O saldo também é privado para garantir que só possa ser
	// modificado através dos métodos (Sacar, Depositar, Transferir).
	saldo float64
}

// O método Sacar para a ContaCorrente.
// A lógica interna é idêntica à da ContaPoupanca, mostrando que tipos
// diferentes podem compartilhar comportamentos iguais.
// O receiver (c *ContaCorrente) é um ponteiro para modificar o saldo.
func (c *ContaCorrente) Sacar(valorDoSaque float64) (string, float64) {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso", c.saldo
	} else {
		return "saldo insuficiente", c.saldo
	}
}

// O método Depositar para a ContaCorrente, também idêntico em lógica
// ao da ContaPoupanca e também usando um ponteiro no receiver para
// modificar o estado da struct.
func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		return "Depósito realizado com sucesso", c.saldo
	} else {
		return "Valor inválido para depósito", c.saldo
	}
}

// Este é um método EXCLUSIVO da ContaCorrente, demonstrando como tipos podem ter
// comportamentos únicos e especializados.
// O receiver é um ponteiro, pois este método modifica o saldo da conta de ORIGEM.
func (c *ContaCorrente) Transferir(valorDaTransferencia float64, contadestino *ContaCorrente) bool {
	// O parâmetro 'contadestino' também é um ponteiro (*ContaCorrente).
	// Isso é CRUCIAL e muito eficiente. Passamos um ponteiro para que o método 'Depositar'
	// que chamaremos a seguir possa modificar a conta de DESTINO original, e não uma cópia dela.

	// Lógica de negócio para a transferência: o valor deve ser positivo e
	// o saldo da conta de origem deve ser suficiente.
	if valorDaTransferencia < c.saldo && valorDaTransferencia > 0 {
		// 1. Debita o valor da conta de origem (a conta 'c').
		c.saldo -= valorDaTransferencia

		// 2. REUTILIZAÇÃO DE CÓDIGO! Em vez de escrever 'contadestino.saldo += ...',
		// nós chamamos o método Depositar da conta de destino. Isso é excelente, pois
		// reutiliza a lógica de validação (valor > 0) que já existe lá, evitando duplicação de código.
		contadestino.Depositar(valorDaTransferencia)

		// O método retorna um booleano simples: 'true' para sucesso.
		return true
	} else {
		// Retorna 'false' para falha.
		return false
	}
}

// O método "getter" para o saldo da ContaCorrente. É idêntico ao da ContaPoupanca
// e serve ao mesmo propósito: expor de forma segura um campo privado.
func (c *ContaCorrente) ObterSaldo() float64 {
	return c.saldo
}