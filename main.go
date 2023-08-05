package main

import (
	"github.com/lucasjct/api-go-gin/database"
	"github.com/lucasjct/api-go-gin/routes"
)

func main() {

	database.ConectaComBancoDeDados()
	routes.HandleRequest()

}
