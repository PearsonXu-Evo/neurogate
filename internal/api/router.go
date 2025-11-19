package api

import (
	"net/http"
	"neurogate/internal/config"

	"github.com/gin-gonic/gin"
)

func NewRouter(cfg *config.Config) *gin.Engine {
	gin.SetMode(cfg.Server.Mode)

	r := gin.New()

	r.Use(gin.Recovery())

	r.Use(gin.Logger())

	r.GET("/healthy", func (c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
			"version": cfg.App.Version,
		})
	})
	
	return r
}
