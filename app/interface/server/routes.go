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
	r.GET("/games", rest.ListGames)
	r.GET("/game/:id", rest.GetGame)
	//r.PATCH("/game/:id", rest.UpdateGame) //this will be used, to pause game or any other action on a game
	r.POST("/game/:id/click", rest.ClickCell)

}
