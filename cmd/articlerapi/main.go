package main

import (
	"context"
	"github.com/Nealoth/articler-api/pkg/articlerapi"
	"github.com/Nealoth/articler-api/pkg/articlerapi/configuration"
	"github.com/Nealoth/articler-api/pkg/articlerapi/logger"
	"github.com/Nealoth/articler-api/pkg/articlerapi/repositories"
	"github.com/Nealoth/articler-api/pkg/articlerapi/routes"
	"github.com/Nealoth/articler-api/pkg/articlerapi/services"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	conf := configuration.ParseConfiguration("./conf/conf.toml")
	ctx, cancel := context.WithCancel(context.Background())

	if !conf.ServerConfiguration.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	server := articlerapi.NewServer(conf)
	db, err := articlerapi.InitPostgresConnection(ctx, conf.DbConfiguration)

	if err != nil {
		log.Fatal(err)
	}

	repoLayer := repositories.InitDbRepositoryLayer(ctx, conf.DbConfiguration, db)
	serviceLayer := services.InitServiceLayer(conf.ServerConfiguration, conf.AuthConfiguration, repoLayer)
	routes.InitRoutes(server, serviceLayer)

	go func() {
		log.Fatal(server.Start(ctx))
	}()

	logger.Info("Server started at: %s", conf.ServerConfiguration.Port)

	osSigChan := make(chan os.Signal, 1)
	signal.Notify(osSigChan, syscall.SIGINT)

	sig := <-osSigChan

	logger.Info("Starting to shut down server. Caught %s signal", sig)

	cancel()
}
