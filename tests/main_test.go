package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gin-gonic/gin"
	"github.com/lucasjct/api-go-gin/controllers"
	"github.com/lucasjct/api-go-gin/database"
	"github.com/lucasjct/api-go-gin/models"
)

var ID int // use this variable in test by id.

func SetupTestRoutes() *gin.Engine {
	gin.SetMode(gin.ReleaseMode) // Simplify the tests output
	routes := gin.Default()
	return routes
}

// Mock Student
func CreateStudentMock() {
	aluno := models.Aluno{Nome: "Nome do Aluno Teste", CPF: "12345678901", RG: "123456789"}
	database.DB.Create(&aluno)
	ID = int(aluno.ID)
}

// Delete Mock Student
func DeleteStudentMock() {
	var aluno models.Aluno
	database.DB.Delete(&aluno, ID)

}

func TesChecktStatusCodeHello(t *testing.T) { // All tests function should be Capitalize and there are 'Test' in function name.

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
	CreateStudentMock()
	defer DeleteStudentMock()
	r := SetupTestRoutes()
	r.GET("/alunos", controllers.ShowAllStudents)
	req, _ := http.NewRequest("GET", "/alunos", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)
	assert.Equal(t, http.StatusOK, response.Code)

}

func TestSearchStudentByCpf(t *testing.T) {
	database.ConectaComBancoDeDados()
	CreateStudentMock()       // evoke mock function for create a student
	defer DeleteStudentMock() // evoke mock function for delete a student
	r := SetupTestRoutes()
	r.GET("/alunos/cpf/:cpf", controllers.SearchByCPF)
	req, _ := http.NewRequest("GET", "/alunos/cpf/12345678901", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)
	assert.Equal(t, http.StatusOK, response.Code)

}

func TestSearchStudentByID(t *testing.T) {

	database.ConectaComBancoDeDados()
	CreateStudentMock()
	defer DeleteStudentMock()
	r := SetupTestRoutes()
	r.GET("/alunos/:id", controllers.SearchStudentById)
	pathOfSearch := "/alunos/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("GET", pathOfSearch, nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)
	var studentMock models.Aluno
	json.Unmarshal(response.Body.Bytes(), &studentMock)
	assert.Equal(t, "Nome do Aluno Teste", studentMock.Nome)
	assert.Equal(t, "12345678901", studentMock.CPF)
	assert.Equal(t, "123456789", studentMock.RG)
	assert.Equal(t, http.StatusOK, response.Code)
}

func TestDeleteStudent(t *testing.T) {
	database.ConectaComBancoDeDados()
	CreateStudentMock()
	r := SetupTestRoutes()
	r.DELETE("/alunos/:id", controllers.DeleteStudent)
	pathOfSearch := "/alunos/" + strconv.Itoa(ID) // get path dynamically
	req, _ := http.NewRequest("DELETE", pathOfSearch, nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)
	assert.Equal(t, http.StatusOK, response.Code)

}
