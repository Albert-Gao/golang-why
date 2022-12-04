package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompute(t *testing.T) {
	nums := []int8{0, 0, 0, 0}

	compute(&nums)

	assert.Equal(t, []int8{1, 1, 1, 1}, nums)
}
