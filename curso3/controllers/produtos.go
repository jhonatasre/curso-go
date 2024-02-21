package controllers

import (
	"log"
	"net/http"
	"strconv"
	"text/template"

	"curso3/models"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	// produtos := []Produto{
	// 	{Nome: "Camiseta", Descricao: "Camiseta preta", Preco: 79.90, Quantidade: 150},
	// 	{Nome: "Calça", Descricao: "Calça jeans", Preco: 99.90, Quantidade: 120},
	// 	{Nome: "Tênis", Descricao: "Tênis esportivo", Preco: 129.90, Quantidade: 55},
	// 	{Nome: "Bermuda", Descricao: "Bermuda de praia", Preco: 49.90, Quantidade: 200},
	// }

	produtos := models.BuscaTodosOsProdutos()

	templates.ExecuteTemplate(w, "Index", produtos)
}

func New(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoConvertido, err := strconv.ParseFloat(preco, 64)

		if err != nil {
			log.Println("Erro na conversão do preço:", err)
			return
		}

		quantidadeConvertida, err := strconv.Atoi(quantidade)

		if err != nil {
			log.Println("Erro na conversão da quantidade:", err)
			return
		}

		models.CriaNovoProduto(nome, descricao, precoConvertido, quantidadeConvertida)
	}

	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	models.DeletaProduto(id)
	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	produto := models.BuscaProduto(id)
	templates.ExecuteTemplate(w, "Edit", produto)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		idConvertido, err := strconv.Atoi(id)

		if err != nil {
			log.Println("Erro na conversão do ID para int:", err)
			return
		}

		precoConvertido, err := strconv.ParseFloat(preco, 64)

		if err != nil {
			log.Println("Erro na conversão do preço:", err)
			return
		}

		quantidadeConvertida, err := strconv.Atoi(quantidade)

		if err != nil {
			log.Println("Erro na conversão da quantidade:", err)
			return
		}

		models.AtualizaProduto(idConvertido, nome, descricao, precoConvertido, quantidadeConvertida)
	}

	http.Redirect(w, r, "/", 301)
}
