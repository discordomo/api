package router

import (
	"github.com/discordomo/api/api"
	"github.com/discordomo/api/router/middleware"
	"github.com/gin-gonic/gin"
)

// ChannelHandlers adds channel handlers to the router group
func ChannelHandlers(base *gin.RouterGroup) {
	channels := base.Group("/channels")
	{
		channels.POST("/", middleware.EstablishSession(), api.CreateChannel)
		channels.DELETE("/:channel", middleware.EstablishSession(), api.DeleteChannel)
	}
}
