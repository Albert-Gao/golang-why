package main

import "time"

func ms(ms int) time.Duration {
	return time.Duration(ms) * time.Millisecond
}

func step(val int) int {
	// time.Sleep(ms(1))
	return val + 1
}

func createArrayWithZeros(length int) []int {
	arr := make([]int, length)

	for i := range arr {
		arr[i] = 0
	}

	return arr
}
