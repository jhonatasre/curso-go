package main

import (
	"curso6/database"
	"curso6/routes"
)

func main() {
	database.ConnectaComBancoDeDados()
	routes.HandleRequests()
}
