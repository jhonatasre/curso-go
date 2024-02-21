package contas

import (
	"fmt"

	"curso2/clientes"
)

type ContaCorrente struct {
	Titular                    clientes.Titular
	NumeroAgencia, NumeroConta int
	saldo                      float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) (bool, float64) {
	if valorDoSaque <= 0 {
		fmt.Println("O valor do saque é inválido")
		return false, c.saldo
	}

	if valorDoSaque <= c.saldo {
		c.saldo -= valorDoSaque
		fmt.Println("Saque realizado com sucesso")
		return true, c.saldo
	}

	fmt.Println("Saldo insuficiente")
	return false, c.saldo
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) (bool, float64) {
	if valorDoDeposito <= 0 {
		fmt.Println("O valor do depósito é inválido")
		return false, c.saldo
	}

	c.saldo += valorDoDeposito
	fmt.Println("Depósito realizado com sucesso")
	return true, c.saldo
}

func (c *ContaCorrente) Transferir(valorDaTransferencia float64, contaDestino *ContaCorrente) (bool, float64) {
	if valorDaTransferencia <= 0 {
		fmt.Println("O valor da transferência é inválido")
		return false, c.saldo
	}

	if valorDaTransferencia <= c.saldo {
		c.saldo -= valorDaTransferencia
		contaDestino.Depositar(valorDaTransferencia)

		fmt.Println("Transferência realizada com sucesso")
		return true, c.saldo
	}

	fmt.Println("Saldo insuficiente")
	return false, c.saldo
}

func (c *ContaCorrente) ObterSaldo() float64 {
	return c.saldo
}
