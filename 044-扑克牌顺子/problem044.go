package main

import (
	"fmt"
	"sort"
)

func poker(nums []int) string {
	sort.Ints(nums)
	start := 0 
	for nums[start] == 0 {
		start++
	}
	for i := start+1; i < len(nums); i++{
		if nums[i] - nums[i-1] -1 > start || nums[i] == nums[i-1] {
			return "Oh my god"
		}
		start -= (nums[i] - nums[i-1] -1)
		fmt.Println(i, start)
	}
	if start >= 0 {
		return "So lucky"
	}
	return "Oh my god"
}

func main() {
	arr1 := []int{ 1, 3, 2, 6, 4 };
    fmt.Println(arr1, " => ", poker(arr1))

    arr2 := []int{ 3, 5, 1, 0, 4, };
    fmt.Println(arr2, " => ", poker(arr2))

    arr3 := []int{ 1, 0, 0, 1, 0 };
    fmt.Println(arr3, " => ", poker(arr3))
}

