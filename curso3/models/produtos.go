package models

import (
	"curso3/db"
)

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func BuscaTodosOsProdutos() []Produto {
	db := db.ConectaBancoDados()

	selectProdutos, err := db.Query("SELECT * FROM produtos order by id asc")

	if err != nil {
		panic(err.Error())
	}

	p := Produto{}

	produtos := []Produto{}

	for selectProdutos.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)

		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)
	}

	defer db.Close()

	return produtos
}

func CriaNovoProduto(nome, descricao string, preco float64, quantidade int) {
	db := db.ConectaBancoDados()

	insertProduto, err := db.Prepare("INSERT INTO produtos(nome, descricao, preco, quantidade) VALUES($1, $2, $3, $4)")

	if err != nil {
		panic(err.Error())
	}

	insertProduto.Exec(nome, descricao, preco, quantidade)

	defer db.Close()
}

func DeletaProduto(id string) {
	db := db.ConectaBancoDados()

	deletarProduto, err := db.Prepare("DELETE FROM produtos WHERE id=$1")

	if err != nil {
		panic(err.Error())
	}

	deletarProduto.Exec(id)

	defer db.Close()
}

func BuscaProduto(id string) Produto {
	db := db.ConectaBancoDados()

	produtoDoBanco, err := db.Query("SELECT * FROM produtos WHERE id=$1", id)

	if err != nil {
		panic(err.Error())
	}

	produtoRetorno := Produto{}

	for produtoDoBanco.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = produtoDoBanco.Scan(&id, &nome, &descricao, &preco, &quantidade)

		if err != nil {
			panic(err.Error())
		}

		produtoRetorno.Id = id
		produtoRetorno.Nome = nome
		produtoRetorno.Descricao = descricao
		produtoRetorno.Preco = preco
		produtoRetorno.Quantidade = quantidade
	}

	defer db.Close()

	return produtoRetorno
}

func AtualizaProduto(id int, nome, descricao string, preco float64, quantidade int) {
	db := db.ConectaBancoDados()

	atualizaProduto, err := db.Prepare("UPDATE produtos SET nome=$1, descricao=$2, preco=$3, quantidade=$4 WHERE id=$5")

	if err != nil {
		panic(err.Error())
	}

	atualizaProduto.Exec(nome, descricao, preco, quantidade, id)

	defer db.Close()
}
