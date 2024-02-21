package routes

import (
	"curso5/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()

	r.GET("/alunos", controllers.ExibeTodosAlunos)
	r.POST("/alunos", controllers.CriaAluno)
	r.GET("/alunos/:id", controllers.ExibeAluno)
	r.PUT("/alunos/:id", controllers.AtualizaAluno)
	r.DELETE("/alunos/:id", controllers.DeletaAluno)
	r.GET("/alunos/cpf/:cpf", controllers.BuscaAlunoPorCPF)

	r.GET("/:nome", controllers.Saudacao)

	r.Run(":8000")
}
