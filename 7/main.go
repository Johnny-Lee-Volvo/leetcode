package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	ret := 0
	for x != 0 {
		ret = ret*10 + x%10
		x /= 10
	}
	if ret <= math.MaxInt32 && ret >= math.MinInt32 {
		return ret
	}
	return 0
}

func main() {
	fmt.Println("123", reverse(123))
	fmt.Println("-123", reverse(-123))
	fmt.Println("120", reverse(120))
	fmt.Println("1534236469", reverse(1534236469))
	fmt.Println("-2147483648", reverse(-2147483648))
}
