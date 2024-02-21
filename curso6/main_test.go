package main

import (
	"bytes"
	"curso6/controllers"
	"curso6/database"
	"curso6/models"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var ID int

func SetupDasRotasDeTestes() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	return r
}

func CriaAlunoMock() {
	aluno := models.Aluno{
		Nome: "João",
		Cpf:  "12345678901",
		Rg:   "123456789",
	}

	database.DB.Create(&aluno)

	ID = int(aluno.Id)
}

func DeletaAlunoMock() {
	var aluno models.Aluno
	database.DB.Delete(&aluno, ID)
}

func TestListandoTodosOsAlunosHandler(t *testing.T) {
	database.ConnectaComBancoDeDados()

	r := SetupDasRotasDeTestes()
	r.GET("/alunos", controllers.ExibeTodosAlunos)

	req, _ := http.NewRequest("GET", "/alunos", nil)
	resp := httptest.NewRecorder()

	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code, "O status code esperado é 200")
}

func TestBuscaAlunoPorCpfHandler(t *testing.T) {
	database.ConnectaComBancoDeDados()
	CriaAlunoMock()

	r := SetupDasRotasDeTestes()
	r.GET("/alunos/cpf/:cpf", controllers.BuscaAlunoPorCPF)

	req, _ := http.NewRequest("GET", "/alunos/cpf/12345678901", nil)
	resp := httptest.NewRecorder()

	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code, "O status code esperado é 200")

	defer DeletaAlunoMock()
}

func TestExibeAlunoHandler(t *testing.T) {
	database.ConnectaComBancoDeDados()
	CriaAlunoMock()

	r := SetupDasRotasDeTestes()
	r.GET("/alunos/:id", controllers.ExibeAluno)

	req, _ := http.NewRequest("GET", "/alunos/"+strconv.Itoa(ID), nil)
	resp := httptest.NewRecorder()

	r.ServeHTTP(resp, req)

	var alunoMock models.Aluno

	json.Unmarshal(resp.Body.Bytes(), &alunoMock)

	assert.Equal(t, "João", alunoMock.Nome, "O nome esperado é João")
	assert.Equal(t, "12345678901", alunoMock.Cpf, "O CPF esperado é 12345678901")
	assert.Equal(t, "123456789", alunoMock.Rg, "O RG esperado é 123456789")

	assert.Equal(t, http.StatusOK, resp.Code, "O status code esperado é 200")

	defer DeletaAlunoMock()
}

func TestDeletaAlunoHandler(t *testing.T) {
	database.ConnectaComBancoDeDados()
	CriaAlunoMock()

	r := SetupDasRotasDeTestes()
	r.DELETE("/alunos/:id", controllers.DeletaAluno)

	req, _ := http.NewRequest("DELETE", "/alunos/"+strconv.Itoa(ID), nil)
	resp := httptest.NewRecorder()

	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code, "O status code esperado é 200")
}

func TestAtualizaAlunoHandler(t *testing.T) {
	database.ConnectaComBancoDeDados()
	CriaAlunoMock()

	r := SetupDasRotasDeTestes()
	r.PUT("/alunos/:id", controllers.AtualizaAluno)

	aluno := models.Aluno{
		Nome: "João",
		Cpf:  "99999999999",
		Rg:   "999999999",
	}
	valorJson, _ := json.Marshal(aluno)

	req, _ := http.NewRequest("PUT", "/alunos/"+strconv.Itoa(ID), bytes.NewBuffer(valorJson))
	resp := httptest.NewRecorder()

	r.ServeHTTP(resp, req)

	var alunoMockAtualizado models.Aluno

	json.Unmarshal(resp.Body.Bytes(), &alunoMockAtualizado)

	assert.Equal(t, "João", alunoMockAtualizado.Nome, "O nome esperado é João")
	assert.Equal(t, "99999999999", alunoMockAtualizado.Cpf, "O CPF esperado é 99999999999")
	assert.Equal(t, "999999999", alunoMockAtualizado.Rg, "O RG esperado é 999999999")

	assert.Equal(t, http.StatusOK, resp.Code, "O status code esperado é 200")

	defer DeletaAlunoMock()
}

// func TestCriaAlunoHandler(t *testing.T) {
// 	database.ConnectaComBancoDeDados()
// 	CriaAlunoMock()

// 	r := SetupDasRotasDeTestes()
// 	r.POST("/alunos", controllers.CriaAluno)

// 	req, _ := http.NewRequest("POST", "/alunos", nil)
// 	resp := httptest.NewRecorder()

// 	r.ServeHTTP(resp, req)

// 	assert.Equal(t, http.StatusBadRequest, resp.Code, "O status code esperado é 400")

// 	defer DeletaAlunoMock()
// }
