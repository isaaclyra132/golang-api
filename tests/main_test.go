package tests

import (
	"api-go/controller"
	"api-go/models"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetFilmes(t *testing.T) {

	filmes := []models.Filme{
		{ID: "1", Titulo: "Hotel Transilvânia", Direcao: "Genndy Tartakovsky, Jennifer Kluska, Derek Drymon", Producao: "Sony Pictures Animation", AnoLancamento: 2012},
		{ID: "2", Titulo: "Tá Dando Onda", Direcao: "Ash Brannon, Chris Buck", Producao: "Sony Pictures Animation", AnoLancamento: 2007},
		{ID: "3", Titulo: "Interestelar", Direcao: "Christopher Nolan", Producao: "Legendary Pictures, Syncopy Films, Lynda Obst Productions", AnoLancamento: 2014},
		{ID: "4", Titulo: "Vingadores: Ultimato", Direcao: "Anthony Russo, Joe Russo", Producao: "Marvel Studios", AnoLancamento: 2019},
		{ID: "5", Titulo: "Coringa", Direcao: "Todd Phillips", Producao: "Village, Roadshow Pictures, DC Films, Sikelia Productions, Joint Effort Productions, Green Hat Films", AnoLancamento: 2019},
	}

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
	// TODO - Testar a rota de criar filme
}

func TestUpdateFilme(t *testing.T) {
	// TODO - Testar a rota de atualizar filme
}

func TestDeleteFilme(t *testing.T) {
	// TODO - Testar a rota de deletar filme
}
