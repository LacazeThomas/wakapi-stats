package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSecondToFormattedTime(t *testing.T) {
	assert.Equal(t, "10mins", SecondToFormattedTime(600))
	assert.Equal(t, "1min", SecondToFormattedTime(60))
	assert.Equal(t, "2mins", SecondToFormattedTime(120))
	assert.Equal(t, "1h10", SecondToFormattedTime(4200))
	assert.Equal(t, "1h01", SecondToFormattedTime(3660))
	assert.Equal(t, "1h02", SecondToFormattedTime(3720))
	assert.Equal(t, "10h", SecondToFormattedTime(3600*10))
	assert.Equal(t, "10h10", SecondToFormattedTime(3600*10+60*10))
	assert.Equal(t, "59mins", SecondToFormattedTime(3540))
}
