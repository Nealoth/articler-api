package v1apiroutes

import (
	"github.com/Nealoth/articler-api/pkg/articlerapi/services"
	"github.com/gin-gonic/gin"
)

type RoutesStore struct {
	serviceLayer *services.ServiceLayer
	basePath     string
	authPath     string
}

const (
	ProfileAPIPath = "/profile"
)

func InitV1UserRoutes(rg *gin.RouterGroup, sl *services.ServiceLayer) *gin.RouterGroup {

	rs := RoutesStore{
		serviceLayer: sl,
		basePath:     rg.BasePath(),
		authPath:     rg.BasePath() + ProfileAPIPath,
	}

	rg.POST("/signUp", rs.signUp)
	rg.POST("/signIn", rs.signIn)

	profileGroup := rg.Group(ProfileAPIPath)
	profileGroup.Use(rs.jwtAuthMiddleware)
	{
		profileGroup.GET("/", rs.userProfile)
		profileGroup.POST("/logout", rs.logout)
	}

	return rg
}
