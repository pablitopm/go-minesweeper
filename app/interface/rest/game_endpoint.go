package rest

import (
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/pablitopm/go-minesweeper/app/domain/model"
	"github.com/pablitopm/go-minesweeper/app/registry"
	"github.com/pablitopm/go-minesweeper/app/usecase"
)

func CreateGame(c *gin.Context) {
	var game model.Game

	err := c.BindJSON(&game)
	if err != nil {
		log.Error("error", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	err = game.Validate()
	if err != nil {
		log.Error("Did not pass validations", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	ctn := c.MustGet("ctn").(*registry.Container)
	useCase := ctn.Resolve("game-usecase").(usecase.GameUsecase)
	game, _ = useCase.StartGame(game)

	c.JSON(http.StatusCreated, game)
}
