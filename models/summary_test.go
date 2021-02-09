package models

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortIntSlice(t *testing.T) {
	s := []SummaryItem{{TotalSeconds: 10}, {TotalSeconds: 10}, {TotalSeconds: 20}, {TotalSeconds: 20}}
	sort.Sort(ItemsSorter(s))
	assert.True(t, sort.IsSorted(ItemsSorter(s)))
}
