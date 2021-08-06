package models

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

var c = []ColorSummaryItem{{SummaryItem{TotalSeconds: 50}, ""}, {SummaryItem{TotalSeconds: 60}, ""}}

func TestSortIntSlice(t *testing.T) {
	sort.Sort(ItemsSorter(c))
	assert.True(t, sort.IsSorted(ItemsSorter(c)))
}

func TestTotalTimeSpent(t *testing.T) {
	assert.Equal(t, 110, TotalTimeSpent(c))
	c[0].TotalSeconds = 0
	assert.Equal(t, 60, TotalTimeSpent(c))
}
