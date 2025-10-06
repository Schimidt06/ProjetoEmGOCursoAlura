// Define que este arquivo pertence ao pacote 'clientes'.
// Este pacote tem a responsabilidade única e clara de lidar com as
// informações relacionadas aos clientes do banco. Manter essa separação
// (clientes, contas, main) é a essência de uma boa arquitetura de software.
package clientes

// Aqui definimos a struct 'Titular'. Ela serve como um modelo ou um 'formulário'
// para guardar as informações de qualquer titular de conta.
type Titular struct {
	// A struct agrupa três campos, todos do tipo 'string'.
	// Para economizar espaço, quando múltiplos campos seguidos são do mesmo tipo,
	// podemos declará-los na mesma linha, informando o tipo apenas uma vez no final.
	Nome, CPF, Profissao string
}