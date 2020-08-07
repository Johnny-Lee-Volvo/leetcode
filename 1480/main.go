package main

import "fmt"

func runningSum(nums []int) []int {
	size := len(nums)
	for i := 1; i < size; i++ {
		nums[i] = nums[i] + nums[i-1]
	}
	return nums
}

func main() {
	fmt.Println("test sum", runningSum([]int{1, 2, 3, 4}))
	fmt.Println("test sum", runningSum([]int{1, 1, 1, 1, 1}))
	fmt.Println("test sum", runningSum([]int{3, 1, 2, 10, 1}))
}