package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRandomTimes99(t *testing.T) {
	m := newMockRandomNumbers()

	m.On("randomMaximumNumber", 727, 1000).Return(100)

	rn := randomTimes99(m)

	assert.Equal(t, 9900, rn)

	m.AssertCalled(t, "randomMaximumNumber", 727, 1000)
}
