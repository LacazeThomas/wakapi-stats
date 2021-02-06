package models

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortIntSlice(t *testing.T) {
	s := []SummaryItem{{"b", 10}, {"a", 10}, {"z", 20}, {"a", 20}}
	sort.Sort(ItemsSorter(s))
	assert.True(t, sort.IsSorted(ItemsSorter(s)))
}
