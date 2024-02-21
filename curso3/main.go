package main

import (
	"net/http"

	"curso3/routes"
)

func main() {
	routes.CarregaRotas()

	http.ListenAndServe(":8000", nil)
}
