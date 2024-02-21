package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"curso4/database"
	"curso4/models"

	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func TodasPersonalidades(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var p []models.Personalidade
	database.DB.Find(&p)

	json.NewEncoder(w).Encode(p)
}

func RetornaUmaPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var personalidade models.Personalidade
	database.DB.First(&personalidade, id)

	json.NewEncoder(w).Encode(personalidade)
}

func CriaPersonalidade(w http.ResponseWriter, r *http.Request) {
	var personalidade models.Personalidade
	json.NewDecoder(r.Body).Decode(&personalidade)

	database.DB.Create(&personalidade)

	json.NewEncoder(w).Encode(personalidade)
}

func RemovePersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	database.DB.Delete(&models.Personalidade{}, id)

	fmt.Fprint(w, "Removido com sucesso")
}

func AtualizaPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var personalidade models.Personalidade
	database.DB.First(&personalidade, id)

	json.NewDecoder(r.Body).Decode(&personalidade)

	database.DB.Save(&personalidade)

	json.NewEncoder(w).Encode(personalidade)
}
