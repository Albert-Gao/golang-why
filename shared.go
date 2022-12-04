package main

import "time"

func ms(ms int) time.Duration {
	return time.Duration(ms) * time.Millisecond
}

func step(val int8) int8 {
	time.Sleep(ms(10))
	return val + 1
}

func createArrayWithZeros(length int8) []int8 {
	arr := make([]int8, length)

	for i := range arr {
		arr[i] = 0
	}

	return arr
}
