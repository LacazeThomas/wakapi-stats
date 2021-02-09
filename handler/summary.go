package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/lacazethomas/wakapi-stats/config"
	"github.com/lacazethomas/wakapi-stats/models"
	"github.com/lacazethomas/wakapi-stats/route"
	"github.com/lacazethomas/wakapi-stats/utils"
)

func HandlerSummary(c *gin.Context) {
	var summaryItems []models.SummaryItems

	appConfig := config.GetAppConfig()

	t := c.Param("type")

	if t == "languages" {
		summaryItems = route.GetSummaryLanguages(appConfig)
	} else {
		summaryItems = route.GetSummaryEditors(appConfig)
	}

	zap.S().Debugf("Receive summary %+v", summaryItems)

	img, err := utils.CreateStatsDiagram(summaryItems)
	utils.Check("creating ", err)

	zap.S().Debugf("Receive images %+v", img)
	c.Data(http.StatusOK, "qr.html", img)
}
