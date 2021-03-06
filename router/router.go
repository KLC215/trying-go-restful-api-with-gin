package router

import (
	"apiserver/handler/sd"
	"apiserver/handler/user"
	"apiserver/router/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {

	g.Use(gin.Recovery())
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)
	g.Use(mw...)

	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "Route not found.")
	})

	g.POST("/login", user.Login)

	u := g.Group("/v1/user")
	u.Use(middleware.AuthMiddleware())
	{
		u.GET("", user.List)
		u.GET("/:username", user.Get)
		u.POST("", user.Create)
		u.PUT("/:id", user.Update)
		u.DELETE("/:id", user.Delete)
	}

	healthCheck := g.Group("/sd")
	{
		healthCheck.GET("/health", sd.HealthCheck)
		healthCheck.GET("/disk", sd.DiskCheck)
		healthCheck.GET("/cpu", sd.CPUCheck)
		healthCheck.GET("/ram", sd.RAMCheck)
	}

	return g
}
