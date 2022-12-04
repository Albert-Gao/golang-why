package main

func compute(nums *[]int) {
	values := *nums

	for i, num := range values {
		values[i] = step(num)
	}
}
