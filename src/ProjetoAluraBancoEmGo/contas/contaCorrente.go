package contas

type ContaCorrente struct {
	Titular       string
	NumeroAgencia int
	NumeroConta   int
	Saldo         float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) (string, float64) {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.Saldo
	if podeSacar {
		c.Saldo -= valorDoSaque
		return "Saque realizado com sucesso", c.Saldo
	} else {
		return "Saldo insuficiente", c.Saldo
	}
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito > 0 {
		c.Saldo += valorDoDeposito
		return "Depósito realizado com sucesso", c.Saldo
	} else {
		return "Valor inválido para depósito", c.Saldo
	}
}

func (c *ContaCorrente) Transferir(valorDaTransferencia float64, contadestino *ContaCorrente) bool {
	if valorDaTransferencia < c.Saldo && valorDaTransferencia > 0 {
		c.Saldo -= valorDaTransferencia
		contadestino.Depositar(valorDaTransferencia)
		return true
	} else {
		return false
	}
}
