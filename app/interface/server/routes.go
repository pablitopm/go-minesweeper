package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pablitopm/go-minesweeper/app/interface/rest"
)

func createRoutes(r *gin.Engine) {
	r.Use(InjectContainer())
	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.POST("/game", rest.CreateGame)

}
