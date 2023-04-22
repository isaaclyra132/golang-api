package router

import (
	"api-go/controller"

	"github.com/gin-gonic/gin"
)

func StartRouter() {

	router := gin.Default()

	router.GET("/filmes", controller.GetFilmes)
	router.GET("/filmes/:id", controller.GetFilme)
	router.POST("/filmes", controller.CreateFilme)
	router.Run("localhost:8080")
}
