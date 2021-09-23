package v1apiroutes

import (
	"github.com/Nealoth/articler-api/pkg/articlerapi/logger"
	"github.com/Nealoth/articler-api/pkg/articlerapi/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (r *RoutesStore) userProfile(c *gin.Context) {
	userID, err := utils.GetUserIDFromContext(c)
	sl := r.serviceLayer

	if err != nil {
		logger.Error("An error occured: %v", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	profile, err := sl.UserService.GetUserProfileByID(userID)
	if err != nil {
		logger.Error("An error occured: %v", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(200, profile)
}
