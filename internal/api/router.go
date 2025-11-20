package api

import (
	"net/http"
	"neurogate/internal/config"
	"neurogate/internal/core"

	"github.com/gin-gonic/gin"
)

func NewRouter(cfg *config.Config, llm core.LLMProvider) *gin.Engine {
	gin.SetMode(cfg.Server.Mode)

	r := gin.New()

	r.Use(gin.Recovery())

	r.Use(gin.Logger())

	// 初始化 Handler
	ChatHandler := NewChatHandler(llm)

	r.GET("/healthy", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"version": cfg.App.Version,
		})
	})

	// API v1 路由组
	v1 := r.Group("/api/v1")
	{
		v1.POST("/chat", ChatHandler.Chat)
	}

	return r
}
