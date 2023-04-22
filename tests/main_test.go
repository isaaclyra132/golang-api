package tests

import (
	"api-go/controller"
	"api-go/db"
	"api-go/models"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetFilmes(t *testing.T) {

	filmes := db.Filmes

	// Cria um router do Gin para testar as rotas
	router := gin.Default()
	router.GET("/filmes", controller.GetFilmes)

	// Cria um request HTTP GET para a rota /filmes
	req, _ := http.NewRequest("GET", "/filmes", nil)

	// Cria um ResponseRecorder para gravar a resposta do servidor
	w := httptest.NewRecorder()

	// Envia a requisição para o servidor
	router.ServeHTTP(w, req)

	// Verific	a se o código de status é 200 OK
	assert.Equal(t, http.StatusOK, w.Code)

	// Decodifica o corpo da resposta para verificar se está igual ao banco de dados inicial
	var response []models.Filme
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.Nil(t, err)
	assert.Equal(t, filmes, response)
}

func TestGetFilme(t *testing.T) {
	//Instancia filme para teste
	filmeTeste := models.Filme{ID: "1", Titulo: "Hotel Transilvânia", Direcao: "Genndy Tartakovsky, Jennifer Kluska, Derek Drymon", Producao: "Sony Pictures Animation", AnoLancamento: 2012}

	// Cria um router do Gin para testar as rotas
	router := gin.Default()
	router.GET("/filmes/:id", controller.GetFilme)

	// Cria um request HTTP GET para a rota /movies/1
	req, _ := http.NewRequest("GET", "/filmes/1", nil)

	// Cria um ResponseRecorder para gravar a resposta do servidor
	w := httptest.NewRecorder()

	// Envia a requisição para o servidor
	router.ServeHTTP(w, req)

	// Verifica se o código de status é 200 OK
	assert.Equal(t, http.StatusOK, w.Code)

	// Decodifica o corpo da resposta para verificar se o filme retornado está correto
	var filme models.Filme
	err := json.Unmarshal(w.Body.Bytes(), &filme)
	assert.Nil(t, err)
	assert.Equal(t, filmeTeste, filme)
}

func TestCreateFilme(t *testing.T) {
	// Cria um request HTTP POST para a rota /movies com um JSON de filme
	jsonStr := `{
		"titulo": "Spider-Man: Into the Spider-Verse",
		"direcao": "Bob Persichetti, Peter Ramsey, Rodney Rothman",
		"producao": "Sony Pictures Animation, Avi Arad, Pascal Pictures, Lord Miller",
		"ano_lancamento": 2019
	 }`
	req, _ := http.NewRequest("POST", "/filmes", bytes.NewBuffer([]byte(jsonStr)))
	req.Header.Set("Content-Type", "application/json")

	// Cria um router do Gin para testar as rotas
	router := gin.Default()
	router.POST("/filmes", controller.CreateFilme)

	// Cria um ResponseRecorder para gravar a resposta do servidor
	w := httptest.NewRecorder()

	// Envia a requisição para o servidor
	router.ServeHTTP(w, req)

	// Verifica se o código de status é 201 Created
	assert.Equal(t, http.StatusCreated, w.Code)

	// Decodifica o corpo da resposta para verificar se o filme criado está correto
	var filme models.Filme
	err := json.Unmarshal(w.Body.Bytes(), &filme)
	assert.Nil(t, err)
	assert.Equal(t, strconv.Itoa(len(db.Filmes)), filme.ID)
	assert.Equal(t, "Spider-Man: Into the Spider-Verse", filme.Titulo)
	assert.Equal(t, "Bob Persichetti, Peter Ramsey, Rodney Rothman", filme.Direcao)
	assert.Equal(t, "Sony Pictures Animation, Avi Arad, Pascal Pictures, Lord Miller", filme.Producao)
	assert.Equal(t, 2019, filme.AnoLancamento)

	// Verifica se o filme foi adicionado ao armazenamento
	assert.Equal(t, filme, db.Filmes[len(db.Filmes)-1])
}

func TestUpdateFilme(t *testing.T) {
	// Cria um request HTTP PUT para a rota /movies/1 com um JSON de filme atualizado
	jsonStr := `{
		"titulo": "Spider-Man: Into the Spider-Verse",
		"direcao": "Bob Persichetti, Peter Ramsey, Rodney Rothman",
		"producao": "Sony Pictures Animation, Avi Arad, Pascal Pictures, Lord Miller",
		"ano_lancamento": 2019
	 }`
	req, _ := http.NewRequest("PUT", "/filmes/1", bytes.NewBuffer([]byte(jsonStr)))
	req.Header.Set("Content-Type", "application/json")

	// Cria um router do Gin para testar as rotas
	router := gin.Default()
	router.PUT("/filmes/:id", controller.UpdateFilme)

	// Cria um ResponseRecorder para gravar a resposta do servidor
	w := httptest.NewRecorder()

	// Envia a requisição para o servidor
	router.ServeHTTP(w, req)

	// Verifica se o código de status é 200 OK
	assert.Equal(t, http.StatusOK, w.Code)

	// Decodifica o corpo da resposta para verificar se o filme atualizado está correto
	var filme models.Filme
	err := json.Unmarshal(w.Body.Bytes(), &filme)
	assert.Nil(t, err)
	assert.Equal(t, models.Filme{ID: "1", Titulo: "Spider-Man: Into the Spider-Verse", Direcao: "Bob Persichetti, Peter Ramsey, Rodney Rothman", Producao: "Sony Pictures Animation, Avi Arad, Pascal Pictures, Lord Miller", AnoLancamento: 2019}, filme)
}

func TestDeleteFilme(t *testing.T) {
	// TODO - Testar a rota de deletar filme

	// Cria um request HTTP DELETE para a rota /movies/1
	req, _ := http.NewRequest("DELETE", "/filmes/1", nil)

	// Cria um router do Gin para testar as rotas
	router := gin.Default()
	router.DELETE("/filmes/:id", controller.DeleteFilme)

	// Cria um ResponseRecorder para gravar a resposta do servidor
	w := httptest.NewRecorder()

	// Envia a requisição para o servidor
	router.ServeHTTP(w, req)

	// Verifica se o código de status é 204 No Content
	assert.Equal(t, http.StatusNoContent, w.Code)

	// Verifica se o segundo filme ainda existe no armazenamento
	assert.Equal(t, models.Filme{ID: "2", Titulo: "Tá Dando Onda", Direcao: "Ash Brannon, Chris Buck", Producao: "Sony Pictures Animation", AnoLancamento: 2007}, db.Filmes[0])
}
