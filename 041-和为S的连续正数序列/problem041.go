package main

import (
	"fmt"
)

func arraySum(target int) []int{
	if target < 3 { return []int{} }
	start, end := 1, 2
	sum := 3
	for end <= target/2+1 {
		if sum == target {
			break
		} else if sum < target {
			end++
			sum += end
		} else {
			sum -= start
			start++
		}
	}
	res := []int{}
	if sum == target {
		for i:=start; i<=end; i++{
			res = append(res, i)
		}
	}
	return res
}

func main() {
	fmt.Println("5 : ", arraySum(5))
	fmt.Println("100 : ", arraySum(100))
}

