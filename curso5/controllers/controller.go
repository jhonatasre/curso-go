package controllers

import (
	"curso5/database"
	"curso5/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Saudacao(c *gin.Context) {
	nome := c.Param("nome")

	c.JSON(200, gin.H{
		"mensagem": "Olá, mundo! " + nome + ", tudo beleza?",
	})
}

// ---

func ExibeTodosAlunos(c *gin.Context) {
	var alunos []models.Aluno

	database.DB.Find(&alunos)

	c.JSON(http.StatusOK, alunos)
}

func ExibeAluno(c *gin.Context) {
	var aluno models.Aluno
	id := c.Param("id")

	if err := database.DB.First(&aluno, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"mensagem": "Aluno não encontrado"})
		return
	}

	c.JSON(http.StatusOK, aluno)
}

func CriaAluno(c *gin.Context) {
	var aluno models.Aluno

	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"mensagem": "Erro ao criar aluno"})
		return
	}

	database.DB.Create(&aluno)
	c.JSON(http.StatusCreated, aluno)
}

func AtualizaAluno(c *gin.Context) {
	var aluno models.Aluno
	id := c.Param("id")

	if err := database.DB.First(&aluno, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"mensagem": "Aluno não encontrado"})
		return
	}

	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"mensagem": "Erro ao atualizar aluno"})
		return
	}

	database.DB.Save(&aluno)
	c.JSON(http.StatusOK, aluno)
}

func DeletaAluno(c *gin.Context) {
	var aluno models.Aluno
	id := c.Param("id")

	if err := database.DB.First(&aluno, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"mensagem": "Aluno não encontrado"})
		return
	}

	database.DB.Delete(&aluno)
	c.JSON(http.StatusNoContent, nil)
}

func BuscaAlunoPorCPF(c *gin.Context) {
	var aluno models.Aluno
	cpf := c.Param("cpf")

	if err := database.DB.Where(&models.Aluno{Cpf: cpf}).First(&aluno).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"mensagem": "Aluno não encontrado"})
		return
	}

	c.JSON(http.StatusOK, aluno)
}
