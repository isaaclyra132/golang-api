package controller

import (
	"api-go/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var filmes = []models.Filme{
	{ID: "1", Titulo: "Hotel Transilvânia", Direcao: "Genndy Tartakovsky, Jennifer Kluska, Derek Drymon", Producao: "Sony Pictures Animation", AnoLancamento: 2012},
	{ID: "2", Titulo: "Tá Dando Onda", Direcao: "Ash Brannon, Chris Buck", Producao: "Sony Pictures Animation", AnoLancamento: 2007},
	{ID: "3", Titulo: "Interestelar", Direcao: "Christopher Nolan", Producao: "Legendary Pictures, Syncopy Films, Lynda Obst Productions", AnoLancamento: 2014},
	{ID: "4", Titulo: "Vingadores: Ultimato", Direcao: "Anthony Russo, Joe Russo", Producao: "Marvel Studios", AnoLancamento: 2019},
	{ID: "5", Titulo: "Coringa", Direcao: "Todd Phillips", Producao: "Village, Roadshow Pictures, DC Films, Sikelia Productions, Joint Effort Productions, Green Hat Films", AnoLancamento: 2019},
}

func GetFilmes(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, filmes)
}

func GetFilme(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de filme inválido"})
		return
	}
	for _, filme := range filmes {
		if filme.ID == strconv.Itoa(id) {
			c.JSON(http.StatusOK, filme)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Filme não encontrado"})
}

func CreateFilme(c *gin.Context) {
	var filme models.Filme
	if err := c.ShouldBindJSON(&filme); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	filme.ID = strconv.Itoa(len(filmes) + 1)
	filmes = append(filmes, filme)
	c.JSON(http.StatusCreated, filme)
}

func UpdateFilme(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de filme inválido"})
		return
	}
	var updatedFilme models.Filme
	if err := c.ShouldBindJSON(&updatedFilme); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for i, filme := range filmes {
		if filme.ID == strconv.Itoa(id) {
			updatedFilme.ID = strconv.Itoa(id)
			filmes[i] = updatedFilme
			c.JSON(http.StatusOK, updatedFilme)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Filme não encontrado"})
}

func DeleteFilme(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de filme inválido"})
		return
	}
	for i, filme := range filmes {
		if filme.ID == strconv.Itoa(id) {
			filmes = append(filmes[:i], filmes[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Filme deletado"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Filme não encontrado"})
}
