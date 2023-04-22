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
	router.PUT("/filmes/:id", controller.UpdateFilme)
	router.DELETE("/filmes/:id", controller.DeleteFilme)

	router.Run(":8043")
}
