package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lucasjct/api-go-gin/controllers"
)

func HandleRequest() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")   // Indetify HTML templates. Pass /template directory.
	r.Static("/assets", "./assets") //Indetify statics files. Pass /assets as a directory.
	r.GET("/alunos", controllers.ShowAllStudents)
	r.GET("/:name", controllers.Hello)
	r.POST("/alunos", controllers.CreateNewStudent)
	r.GET("/alunos/:id", controllers.SearchStudentById)
	r.DELETE("/alunos/:id", controllers.DeleteStudent)
	r.PATCH("/alunos/:id", controllers.UpdateStudent)
	r.GET("/alunos/cpf/:cpf", controllers.SearchByCPF)
	r.GET("/index", controllers.IndexHTML)
	r.NoRoute(controllers.RouteNotFound) // use when the return is not found
	r.Run()

}
