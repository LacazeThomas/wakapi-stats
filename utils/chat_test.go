package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/lacazethomas/wakapi-stats/models"
)

func TestCreateStatsDiagram(t *testing.T) {
	summaryItems := []models.SummaryItem{}
	summaryItem := models.SummaryItem{
		Digital:      "2:5:25",
		Hours:        2,
		Minutes:      5,
		Name:         "macOS",
		Percent:      6.41,
		Seconds:      25,
		Text:         "2 hrs 5 mins",
		TotalSeconds: 7525,
	}
	summaryItems = append(summaryItems, summaryItem)

	_, err := CreateStatsDiagram(summaryItems)
	assert.ErrorIs(t, nil, err)

	_, err = CreateStatsDiagram(nil)
	assert.Error(t, err)
	assert.Regexp(t, err.Error(), "This item is empty")

}
