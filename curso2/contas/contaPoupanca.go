package contas

import (
	"fmt"

	"curso2/clientes"
)

type ContaPoupanca struct {
	Titular                              clientes.Titular
	NumeroAgencia, NumeroConta, Operacao int
	saldo                                float64
}

func (c *ContaPoupanca) Sacar(valorDoSaque float64) (bool, float64) {
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

func (c *ContaPoupanca) Depositar(valorDoDeposito float64) (bool, float64) {
	if valorDoDeposito <= 0 {
		fmt.Println("O valor do depósito é inválido")
		return false, c.saldo
	}

	c.saldo += valorDoDeposito
	fmt.Println("Depósito realizado com sucesso")
	return true, c.saldo
}

func (c *ContaPoupanca) ObterSaldo() float64 {
	return c.saldo
}
