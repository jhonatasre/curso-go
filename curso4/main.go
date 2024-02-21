package main

import (
	"fmt"

	"curso4/database"
	"curso4/models"
	"curso4/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Jhonatas Reis Matos", Historia: "Desenvolvedor de software"},
		{Id: 2, Nome: "Steve Jobs", Historia: "Fundador da Apple"},
		{Id: 3, Nome: "Bill Gates", Historia: "Fundador da Microsoft"},
		{Id: 4, Nome: "Elon Musk", Historia: "Fundador da Tesla"},
	}

	database.ConnectaComBancoDeDados()
	fmt.Println("Iniciando o servidor Rest com Go!")
	routes.HandlerRequest()
}
