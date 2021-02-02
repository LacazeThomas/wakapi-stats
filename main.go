package main

import (
	"log"

	"go.uber.org/zap"

	"github.com/lacazethomas/wakapi-stats/config"
	"github.com/lacazethomas/wakapi-stats/models"
	"github.com/lacazethomas/wakapi-stats/route"
	"github.com/lacazethomas/wakapi-stats/utils"
)

func main() {
	var logger *zap.Logger
	var err error

	appConfig, gitConfig := config.Load()

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

	summary := route.GetSummary(appConfig)

	zap.S().Debugf("Receive summary %+v", summary)

	for key, element := range neededStats(summary) {

		err = utils.CreateStatsDiagram(element, key)
		utils.Check("creating "+key, err)
		err = utils.PublishToGithub(gitConfig, key)
		utils.Check("publishing "+key, err)

	}
}

var neededStats = func(summary models.Summary) map[string][]models.SummaryItems {

	return map[string][]models.SummaryItems{
		"editors.png": summary.Editors, "languages.png": summary.Languages, "operatingSystems.png": summary.OperatingSystems,
	}
}
