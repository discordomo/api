package router

import (
	"github.com/discordomo/api/api"
	"github.com/discordomo/api/router/middleware"
	"github.com/gin-gonic/gin"
)

// RoleHandlers adds channel handlers to the router group
func RoleHandlers(base *gin.RouterGroup) {
	roles := base.Group("/roles")
	{
		roles.POST("/", middleware.EstablishSession(), api.CreateRole)
		roles.DELETE("/:role", middleware.EstablishSession(), api.DeleteRole)
	}
}
