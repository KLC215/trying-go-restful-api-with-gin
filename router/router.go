package router

import (
	"apiserver/handler/sd"
	"apiserver/router/middleware"

	"github.com/gin-gonic/gin"
)

func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {

	g.Use(gin.Recovery())
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)

	healthCheck := g.Group("/sd")
	{
		healthCheck.GET("/health", sd.HealthCheck)
		healthCheck.GET("/disk", sd.DiskCheck)
		healthCheck.GET("/cpu", sd.CPUCheck)
		healthCheck.GET("/ram", sd.RAMCheck)
	}

	return g
}
