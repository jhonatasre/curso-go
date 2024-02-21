package database

import (
	"curso6/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectaComBancoDeDados() {
	stringDeConexao := "host=localhost user=postgres password=admin dbname=alunos port=15432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(stringDeConexao))

	if err != nil {
		log.Panic("Erro ao conectar com o banco de dados")
	}

	DB.AutoMigrate(&models.Aluno{})
}
