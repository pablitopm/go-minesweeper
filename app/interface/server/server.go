package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	createRoutes(r)
	return r
}

// CreateServer starts a server in port 8080
func CreateServer() *http.Server {
	r := setupRouter()
	return &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
}
