package controller

import (
	"api-go/db"
	"api-go/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetFilmes(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, db.Filmes)
}

func GetFilme(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de filme inválido"})
		return
	}
	for _, filme := range db.Filmes {
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
	filme.ID = strconv.Itoa(len(db.Filmes) + 1)
	db.Filmes = append(db.Filmes, filme)
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
	for i, filme := range db.Filmes {
		if filme.ID == strconv.Itoa(id) {
			updatedFilme.ID = strconv.Itoa(id)
			db.Filmes[i] = updatedFilme
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
	for i, filme := range db.Filmes {
		if filme.ID == strconv.Itoa(id) {
			db.Filmes = append(db.Filmes[:i], db.Filmes[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Filme deletado"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Filme não encontrado"})
}
