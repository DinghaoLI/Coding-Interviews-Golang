package main

import (
	"fmt"
)

func again(nums []int) int {
	appear := make(map[int]bool)
	for _, v := range nums {
		if _, ok := appear[v]; ok {
			return v
		}
		appear[v] = true
	}
	// Not Found Error
	return len(nums) + 1

}

func main() {

	list := []int{2, 3, 1, 0, 2, 5, 3}
	fmt.Println(list, " => ", again(list))
	list = []int{2, 1, 3, 1, 4}
	fmt.Println(list, " => ", again(list))

}
