package main

import (
	"github.com/lucasjct/api-go-gin/database"
	"github.com/lucasjct/api-go-gin/models"
	"github.com/lucasjct/api-go-gin/routes"
)

func main() {

	database.ConectaComBancoDeDados()
	models.Alunos = []models.Aluno{

		{Nome: "teste", CPF: "998544444", RG: "88585888"},
		{Nome: "NAEM teste", CPF: "11112255", RG: "8955555888"},
	}
	routes.HandleRequest()

}
