package routes

import (
	"fmt"
	"github.com/Nealoth/articler-api/pkg/articlerapi"
	"github.com/Nealoth/articler-api/pkg/articlerapi/logger"
	v1_Routes "github.com/Nealoth/articler-api/pkg/articlerapi/routes/v1"
	"github.com/Nealoth/articler-api/pkg/articlerapi/services"
	"github.com/gin-gonic/gin"
)

type APIVersion string
type routeInitializerFunc func(r *gin.RouterGroup, sl *services.ServiceLayer) *gin.RouterGroup

var initializersStorage = map[string]routeInitializerFunc{
	"v1": v1_Routes.InitV1UserRoutes,
}

func InitRoutes(server *articlerapi.Server, sl *services.ServiceLayer) {
	r := server.Engine

	apiVersion := server.Conf.ServerConfiguration.APIVersion

	initFunc, initFuncFound := initializersStorage[apiVersion]

	if initFuncFound {
		rGroup := r.Group(fmt.Sprintf("/api/%s", apiVersion))
		initFunc(rGroup, sl)
		logger.Debug("Initialized routes group: %s", apiVersion)
	} else {
		logger.Warn("Cannot find any route initializers for group: %s", apiVersion)
	}
}
