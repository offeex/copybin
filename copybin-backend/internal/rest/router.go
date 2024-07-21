package rest

import (
	"copybin/internal/config"
	"copybin/internal/rest/handler"
	"github.com/gin-gonic/gin"
)

func NewRouter(cfg *config.Config) (*gin.Engine, error) {
	router := gin.Default()
	restHandler := handler.NewHandler(cfg)

	router.GET("/me", restHandler.Me)

	pastas := router.Group("/pastas")
	pastas.Use(restHandler.AuthMiddleware)
	{
		pastas.GET("/", restHandler.GetPastas)
		pastas.POST("/", restHandler.CreatePasta)
	}

	authDiscord := router.Group("/auth/discord")
	{
		authDiscord.GET("/login", restHandler.DiscordLogin)
		authDiscord.GET("/callback", restHandler.DiscordCallback)
	}
	return router, nil
}
