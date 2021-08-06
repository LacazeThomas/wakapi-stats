package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/lacazethomas/wakapi-stats/models"
)

func TestCreateStatsDiagram(t *testing.T) {
	colorSummaryItems := []models.ColorSummaryItem{{models.SummaryItem{TotalSeconds: 50, Name: "Go"}, "#112233"}, {models.SummaryItem{TotalSeconds: 60, Name: "John"}, ""}}

	_, err := CreateStatsDiagram(colorSummaryItems, false)
	assert.ErrorIs(t, nil, err)

	_, err = CreateStatsDiagram(colorSummaryItems, true)
	assert.ErrorIs(t, nil, err)

	_, err = CreateStatsDiagram(nil, false)
	assert.Error(t, err)
	assert.Regexp(t, err.Error(), "This item is empty")
}

func TestCreatePieFromChartValues(t *testing.T) {
	_, err := createPieFromChartValues(nil)
	assert.Error(t, err)
	assert.Regexp(t, err.Error(), "error rendering image from graph: please provide at least one value\" to match ")
}
