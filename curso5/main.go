package main

import (
	"curso5/database"
	"curso5/routes"
)

func main() {
	database.ConnectaComBancoDeDados()
	routes.HandleRequests()
}
