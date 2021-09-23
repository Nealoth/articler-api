package v1apiroutes

import (
	"github.com/Nealoth/articler-api/pkg/articlerapi/logger"
	"github.com/gin-gonic/gin"
)

//TODO replace
func respondError(c *gin.Context, status int, err error) {
	logger.Debug("[AuthMiddleware] Auth failed with status %s: %s", status, err)
	c.AbortWithStatus(status)
}
