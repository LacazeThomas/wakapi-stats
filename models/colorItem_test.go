package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestColorSummaryItems(t *testing.T) {
	summaryItems := []SummaryItem{{TotalSeconds: 50, Name: "Go"}}

	c, err := ColorSummaryItems(summaryItems, "../colors.json")
	assert.ErrorIs(t, nil, err)
	assert.Equal(t, "#00ADD8", c[0].Color)

}
