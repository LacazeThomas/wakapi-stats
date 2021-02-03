package models

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func TestSortIntSlice(t *testing.T) {
	s := []SummaryItems{SummaryItems{"b",10},SummaryItems{"a",10},SummaryItems{"z",20},SummaryItems{"a",20}}
	sort.Sort(ItemsSorter(s))
	assert.True(t, sort.IsSorted(ItemsSorter(s)))
}

