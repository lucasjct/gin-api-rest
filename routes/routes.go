package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lucasjct/api-go-gin/controllers"
)

func HandleRequest() {

	r := gin.Default()
	r.GET("/alunos", controllers.ShowAllAlumns)
	r.GET("/:name", controllers.Hello)
	r.POST("/alunos", controllers.CreateNewStudent)
	r.Run(":5000")

}
