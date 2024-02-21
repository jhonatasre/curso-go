package main

import (
	"fmt"

	"curso2/contas"
)

func main() {
	contaPoupanca := contas.ContaPoupanca{}
	contaPoupanca.Depositar(100)
	PagarBoleto(&contaPoupanca, 60)

	fmt.Println(contaPoupanca.ObterSaldo())

	// clienteJhonatas := clientes.Titular{Nome: "Jhonatas", CPF: "123.123.123-12", Profissao: "Desenvolvedor"}

	// contaJhonatas := contas.ContaCorrente{
	// 	Titular:       clienteJhonatas,
	// 	NumeroAgencia: 123,
	// 	NumeroConta:   12346,
	// 	Saldo:         100,
	// }

	// fmt.Println(contaJhonatas)

	// contaJhonatas := ContaCorrente{}
	// contaJhonatas.titular = "Jhonatas"
	// contaJhonatas.saldo = 500

	// fmt.Println(contaJhonatas.saldo)

	// contaJhonatas.Sacar(100)
	// contaJhonatas.Depositar(300)
	// contaJhonatas.Depositar(200)
	// contaJhonatas.Depositar(100)

	// fmt.Println(contaJhonatas.saldo)

	// contaJhonatas := contas.ContaCorrente{Titular: "Jhonatas", Saldo: 500}
	// contaDuda := contas.ContaCorrente{Titular: "Duda", Saldo: 1000}

	// fmt.Println(contaJhonatas)
	// fmt.Println(contaDuda)

	// contaJhonatas.Transferir(200, &contaDuda)

	// fmt.Println(contaJhonatas)
	// fmt.Println(contaDuda)
}

func PagarBoleto(conta verificarConta, valorBoleto float64) {
	conta.Sacar(valorBoleto)
}

type verificarConta interface {
	Sacar(valor float64) (bool, float64)
}
