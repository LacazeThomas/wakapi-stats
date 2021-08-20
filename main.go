package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thinkerou/favicon"
	"go.uber.org/zap"

	"github.com/lacazethomas/wakapi-stats/config"
	"github.com/lacazethomas/wakapi-stats/handler"
)

func main() {
	var logger *zap.Logger
	var err error

	gin.SetMode(gin.ReleaseMode)

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
	r.Use(favicon.New("./favicon.ico"))
	r.GET("/api/v1/:type", handler.Summary)
	r.GET("/", func(c *gin.Context) {
		c.Data(http.StatusOK, "text/html; charset=utf-8", []byte("Hi 👋, this is an <a href=\"/api/v1/languages?url=https://stats-code.thomaslacaze.fr/api/v1/users/thomaslacaze/stats/30_days\">example</a>"))
	})

	err = r.Run(":8080")
	if err != nil {
		log.Println("Unable to start web server")
	}

}
