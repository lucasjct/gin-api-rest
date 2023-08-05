package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lucasjct/api-go-gin/controllers"
)

func HandleRequest() {
	r := gin.Default()
	r.GET("/alunos", controllers.ShowAllStudents)
	r.GET("/:name", controllers.Hello)
	r.POST("/alunos", controllers.CreateNewStudent)
	r.GET("/alunos/:id", controllers.SearchStudentById)
	r.DELETE("/alunos/:id", controllers.DeleteStudent)
	r.PATCH("/alunos/:id", controllers.UpdateStudent)
	r.Run()

}
