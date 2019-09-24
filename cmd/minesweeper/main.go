package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/pablitopm/go-minesweeper/app/interface/server"
)

func main() {
	log.Info("Starting API")
	server.StartServer()
}
