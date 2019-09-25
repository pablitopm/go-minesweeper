package server

import (
	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/pablitopm/go-minesweeper/app/registry"
)

func InjectContainer() gin.HandlerFunc {
	log.Debug("Starting container")
	ctn, err := registry.NewContainer()
	if err != nil {
		log.Fatalf("failed to build container: %v", err)
	}
	return func(c *gin.Context) {
		c.Set("ctn", ctn)
		c.Next()
	}
}
