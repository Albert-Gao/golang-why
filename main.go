package main

import (
	"time"
)

func ms(ms int) time.Duration {
	return time.Duration(ms) * time.Millisecond
}

func step(val int8) int8 {
	time.Sleep(ms(100))
	return val + 1
}

func compute(nums *[]int8) {
	values := *nums

	for i, num := range values {
		values[i] = step(num)
	}
}
