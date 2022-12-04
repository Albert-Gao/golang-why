package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestComputeConcurrent(t *testing.T) {
	nums := []int8{0, 1, 2, 3}

	computeConcurrent(&nums)

	assert.Equal(t, []int8{1, 2, 3, 4}, nums)
}

func BenchmarkComputeConcurrent(b *testing.B) {
	nums := createArrayWithZeros(100)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		computeConcurrent(&nums)
	}
}
