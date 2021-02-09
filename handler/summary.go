package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
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

	if t == "" || appURI == "" {
		checkError(c, errors.New("Missing param or url"))
	}

	summary, err := route.GetSummary(appURI)
	checkError(c, err)

	summaryItem = availableStats(summary)[t]
	zap.S().Debugf("Receive summary %+v", summaryItem)

	result, err = utils.CreateStatsDiagram(summaryItem)
	checkError(c, err)

	zap.S().Debugf("Receive images %+v", result)
	c.Data(http.StatusOK, "image/png", result)
}

var availableStats = func(summary models.Summary) map[string][]models.SummaryItem {

	return map[string][]models.SummaryItem{
		"editors": summary.Data.Editors, "languages": summary.Data.Languages, "operatingSystems": summary.Data.OperatingSystems, "machines": summary.Data.Machines, "projects": summary.Data.Projects,
	}
}

func checkError(c *gin.Context, err error) {
	if err != nil {
		zap.S().Error(err)
		c.Error(err)
		c.AbortWithStatusJSON(200, gin.H{"message": err.Error()})
	}
}
