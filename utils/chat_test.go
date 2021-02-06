package utils

import (
	"testing"

	"github.com/lacazethomas/wakapi-stats/models"
	"github.com/stretchr/testify/assert"
)

func TestCreateStatsDiagram(t *testing.T) {
	summaries := models.Summary{
		UserID: "thomaslacaze",
		From:   struct{}{},
		To:     struct{}{},
		Projects: []models.SummaryItem{
			{
				Key:   "space",
				Total: 73012,
			},
		},
		Languages:        []models.SummaryItem{},
		Editors:          []models.SummaryItem{},
		OperatingSystems: []models.SummaryItem{},
		Machines:         []models.SummaryItem{},
	}

	err := CreateStatsDiagram(summaries.Projects, "projects.png")
	assert.ErrorIs(t, nil, err)
	assert.FileExists(t, "projects.png")

	err = CreateStatsDiagram(summaries.Projects, "null/projects.png")
	assert.Error(t, err)
	assert.Regexp(t, ".*error creating image: .*", err.Error())
	assert.NoFileExists(t, "null/projects.png")
}
