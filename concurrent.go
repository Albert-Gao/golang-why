package main

import "sync"

type result struct {
	index  int
	result int
}

func computeConcurrent(nums *[]int) {
	values := *nums
	resultChannel := make(chan result)

	for i, num := range values {
		go func(index int, n int) {
			resultChannel <- result{index, step(n)}
		}(i, num)
	}

	for j := 0; j < len(values); j++ {
		r := <-resultChannel
		values[r.index] = r.result
	}
}

func computeConcurrent2(nums *[]int) {
	var wg sync.WaitGroup
	values := *nums

	for i, num := range values {
		wg.Add(1)
		go func(index int, n int, v []int) {
			defer wg.Done()
			v[index] = step(n)
		}(i, num, values)
	}

	wg.Wait()
}
