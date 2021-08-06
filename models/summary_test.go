package models

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortIntSlice(t *testing.T) {

	c := []ColorSummaryItem{{SummaryItem{TotalSeconds: 50}, ""}, {SummaryItem{TotalSeconds: 60}, ""}}
	sort.Sort(ItemsSorter(c))
	assert.True(t, sort.IsSorted(ItemsSorter(c)))
}
