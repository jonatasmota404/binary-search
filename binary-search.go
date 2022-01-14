package main

import "fmt"

func main() {
	fmt.Println(generateList(10))
}

func generateList(x int) []int {
	nums := []int{}
	for i := 0; i < x; i++ {
		nums = append(nums, i)
	}

	return nums
}
