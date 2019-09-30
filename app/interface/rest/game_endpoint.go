package rest

import (
	"fmt"
	"net/http"
	"strconv"

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
		log.Error("error", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	err = game.Validate()
	if err != nil {
		log.Error("Did not pass validations", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	ctn := c.MustGet("ctn").(*registry.Container)
	useCase := ctn.Resolve("game-usecase").(usecase.GameUsecase)
	game, _ = useCase.StartGame(game)

	c.JSON(http.StatusCreated, game)
}

func ListGames(c *gin.Context) {
	ctn := c.MustGet("ctn").(*registry.Container)
	useCase := ctn.Resolve("game-usecase").(usecase.GameUsecase)
	games, _ := useCase.ListGames()

	c.JSON(http.StatusOK, games)
}

func GetGame(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Param("id"))
	ctn := c.MustGet("ctn").(*registry.Container)
	useCase := ctn.Resolve("game-usecase").(usecase.GameUsecase)
	game, err := useCase.GetGame(ID)
	if err != nil {
		fmt.Println(err)
		c.String(http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, game)
}

func ClickCell(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Param("id"))
	req := struct {
		Col int `json:"col"`
		Row int `json:"row"`
	}{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	ctn := c.MustGet("ctn").(*registry.Container)
	useCase := ctn.Resolve("game-usecase").(usecase.GameUsecase)

	game, err := useCase.GetGame(ID)
	if err != nil {
		log.Errorf("error: could not find Game with ID %d", ID)
		c.String(http.StatusNotFound, "error: could not find Game")
		return
	}

	game, err = useCase.ClickCell(game, req.Row, req.Col)

	if err != nil {
		log.Error(err.Error())
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, game)
}
