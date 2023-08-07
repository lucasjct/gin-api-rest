package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gin-gonic/gin"
	"github.com/lucasjct/api-go-gin/controllers"
	"github.com/lucasjct/api-go-gin/database"
)

func SetupTestRoutes() *gin.Engine {
	routes := gin.Default()
	return routes
}

func TesChecktStatuCodeHello(t *testing.T) { // All tests function should be Capitalize and there are 'Test' in function name.

	r := SetupTestRoutes()             // instance of routes.
	r.GET("/:nome", controllers.Hello) // endpoint that will be tested.
	req, _ := http.NewRequest("GET", "/anyname", nil)
	response := httptest.NewRecorder() // this function receive all response request and storage in variable.
	r.ServeHTTP(response, req)         // this function work doing request. Your resposability is make requests.
	assert.Equal(t, http.StatusOK, response.Code, "Should be equal")
	mockResponse := `{"API diz":"Olá "}`                // mock response body.
	responseBody, _ := ioutil.ReadAll(response.Body)    //get response body.
	assert.Equal(t, mockResponse, string(responseBody)) // assertion with testfy.
}

func TestCheckShowAllStudentsFromList(t *testing.T) {
	database.ConectaComBancoDeDados() // database connection.
	r := SetupTestRoutes()
	r.GET("/alunos", controllers.ShowAllStudents)
	req, _ := http.NewRequest("GET", "/alunos", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)
	assert.Equal(t, http.StatusOK, response.Code)

}
