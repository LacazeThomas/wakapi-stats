package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"go.uber.org/zap"

	"github.com/lacazethomas/wakapi-stats/models"
	"github.com/lacazethomas/wakapi-stats/route"
	"github.com/lacazethomas/wakapi-stats/utils"
)

//Summary return corresponding image
func Summary(c *gin.Context) {
	var summaryItem []models.SummaryItem
	var result []byte

	t := c.Param("type")

	timeSpent, err := strconv.ParseBool(c.DefaultQuery("timeSpent", "false"))
	checkError(c, err)

	appURI := c.Query("url")

	zap.S().Debugf("Receive  %+v %+v", t, appURI)
	if t == "" || appURI == "" {
		checkError(c, errors.New("Missing param or url"))
	}

	summary, err := route.GetSummary(appURI)
	checkError(c, err)

	summaryItem = availableStats(summary)[t]
	zap.S().Debugf("Receive summary %+v", summaryItem)

	colorSummaryItem, err := models.ColorSummaryItems(summaryItem, "colors.json")
	checkError(c, err)

	result, err = utils.CreateStatsDiagram(colorSummaryItem, timeSpent)
	checkError(c, err)

	c.Data(http.StatusOK, "image/svg+xml", result)
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
