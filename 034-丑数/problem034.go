package main

import (
	"fmt"
)

func getUgly(N int) []int {
	res := make([]int, N)
	if N < 1 { return res }
	res[0] = 1
	index2, index3, index5 := 0, 0, 0
	index := 1
	for index <= N-1 { 
		minValue := min(min(res[index2]*2, res[index3]*3), res[index5]*5)
		if minValue == res[index2]*2{
			index2++
		}
		// 不是elseif 因为可能重复
		if minValue == res[index3]*3{
			index3++
		}
		if minValue == res[index5]*5{
			index5++
		}
		res[index] = minValue
		index++
	}
	return res
}

func min(a,b int) int {
	if a<b {
		return a
	}
	return b
}

func main() {
	fmt.Println(getUgly(50))
}
















