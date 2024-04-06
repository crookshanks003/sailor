package server

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/pritesh-mantri/sailor/config"
	"github.com/pritesh-mantri/sailor/internal/data"
)

type Application struct {
	Config config.Config
	Models data.Models
}

func (app *Application) Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func (app *Application) Routes() *gin.Engine {
	r := gin.Default()
	api := r.Group("/api")
	{
		api.GET("/health", app.Health)
	}

	return r
}
