package main

type result struct {
	index  int
	result int8
}

func computeConcurrent(nums *[]int8) {
	values := *nums
	resultChannel := make(chan result)

	for i, num := range values {
		go func(index int, n int8) {
			resultChannel <- result{index, step(n)}
		}(i, num)
	}

	for j := 0; j < len(values); j++ {
		r := <-resultChannel
		values[r.index] = r.result
	}
}
