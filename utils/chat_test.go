package utils

import (
	"bytes"
	"github.com/lacazethomas/wakapi-stats/models"
	"github.com/wcharczuk/go-chart"
	"testing"
)

func TestCreateStatsDiagram(t *testing.T) {
	// replaced new assertions helper



	chartValue := {chartValue2,chartValue1}


	CreateStatsDiagram(summaryItems []models.SummaryItems, imageName string)

	b := bytes.NewBuffer([]byte{})
	pie.Render(PNG, b)
	testutil.AssertNotZero(t, b.Len())
}