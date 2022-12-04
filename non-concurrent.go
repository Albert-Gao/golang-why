package main

func compute(nums *[]int8) {
	values := *nums

	for i, num := range values {
		values[i] = step(num)
	}
}
