package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompute(t *testing.T) {
	nums := []int{0, 1, 2, 3}

	compute(&nums)

	assert.Equal(t, []int{1, 2, 3, 4}, nums)
}

func TestComputeConcurrent(t *testing.T) {
	nums := []int{0, 1, 2, 3}

	computeConcurrent(&nums)

	assert.Equal(t, []int{1, 2, 3, 4}, nums)
}

func TestComputeConcurrent2(t *testing.T) {
	nums := []int{0, 1, 2, 3}

	computeConcurrent2(&nums)

	assert.Equal(t, []int{1, 2, 3, 4}, nums)
}

func BenchmarkCompute(b *testing.B) {
	nums := createArrayWithZeros(100)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		compute(&nums)
	}
}

func BenchmarkComputeConcurrent(b *testing.B) {
	nums := createArrayWithZeros(100)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		computeConcurrent(&nums)
	}
}

func BenchmarkComputeConcurrent2(b *testing.B) {
	nums := createArrayWithZeros(100)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		computeConcurrent2(&nums)
	}
}
