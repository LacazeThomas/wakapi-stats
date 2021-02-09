package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/lacazethomas/wakapi-stats/models"
	"github.com/lacazethomas/wakapi-stats/route"
	"github.com/lacazethomas/wakapi-stats/utils"
)

func HandlerSummary(c *gin.Context) {
	var summaryItem []models.SummaryItem
	var result []byte

	t := c.Param("type")
	appURI := c.Query("url")

	if len(t) == 0 || len(appURI) == 0 {
		zap.S().Errorf("Missing param or query into your url")
		result = []byte("Error has been detected, see logs")
		c.Data(http.StatusOK, "text/plain", result)
	}

	summary, err := route.GetSummary(appURI)
	if err != nil {
		zap.S().Errorf("http error %s", err.Error())
		result = []byte("Error has been detected, see logs")
		c.Data(http.StatusOK, "text/plain", result)
	}

	summaryItem = availableStats(summary)[t]
	zap.S().Debugf("Receive summary %+v", summaryItem)

	result, err = utils.CreateStatsDiagram(summaryItem)
	if err != nil {
		zap.S().Errorf("creating %s", err.Error())
		result = []byte("Error has been detected, see logs")
	}

	zap.S().Debugf("Receive images %+v", result)
	c.Data(http.StatusOK, "text/plain", result)
}

var availableStats = func(summary models.Summary) map[string][]models.SummaryItem {

	return map[string][]models.SummaryItem{
		"editors": summary.Data.Editors, "languages": summary.Data.Languages, "operatingSystems": summary.Data.OperatingSystems, "machines": summary.Data.Machines, "projects": summary.Data.Projects,
	}
}
