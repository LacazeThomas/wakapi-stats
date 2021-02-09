package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/lacazethomas/wakapi-stats/config"
	"github.com/lacazethomas/wakapi-stats/handler"
)

func main() {
	var logger *zap.Logger
	var err error

	appConfig := config.GetAppConfig()

	// Set log level
	if appConfig.IsDev() {
		logger, err = zap.NewDevelopment()
	} else {
		logger, err = zap.NewProduction()
	}
	if err != nil {
		log.Println("Error to initialize logger")
	}

	defer logger.Sync()
	zap.ReplaceGlobals(logger)

	r := gin.Default()
	r.GET("/stats/:type", handler.HandlerSummary)
	r.Run(":8080")
}
