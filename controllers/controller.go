package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucasjct/api-go-gin/database"
	"github.com/lucasjct/api-go-gin/models"
)

func ShowAllAlumns(c *gin.Context) {
	c.JSON(200, models.Alunos)
}

// "c" property is about the request's context. So, we can acces some features when type "c."
// c property get all context, he has all controll over request's context.
func Hello(c *gin.Context) {

	// Params's gin. Is the resource from Context. So with c poperty we can define the url's params.
	nome := c.Params.ByName("nome")
	c.JSON(200, gin.H{
		"API diz": "Olá " + nome,
	})
}

func CreateNewStudent(c *gin.Context) {
	var aluno models.Aluno
	if err := c.ShouldBindJSON(&aluno); err != nil { // ShouldBindJSON - get all response body from request
		c.JSON(http.StatusBadRequest, gin.H{
			"errr": err.Error()})
		return

	}
	database.DB.Create(&aluno)   // create data
	c.JSON(http.StatusOK, aluno) // feedback request ok
}
