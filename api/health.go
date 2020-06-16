package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Health returns a 200, ok
func Health(c *gin.Context) {
	c.JSON(http.StatusOK, "ok")
}
