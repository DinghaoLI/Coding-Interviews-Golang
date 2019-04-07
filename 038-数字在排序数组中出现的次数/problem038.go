package main

import (
	"fmt"
)

// Binary search
func count(nums []int, target int) int {

	left, right := 0, len(nums)-1
	var mid int
	for left < right {
		mid = (left + right) / 2
		if nums[mid] == target {
			for nums[mid] != nums[right] {
				right--
			}
			for nums[mid] != nums[left] {
				left++
			}
			break
		}
		if nums[mid] > target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if left < right {
		return right-left+1
	}
	return -1
}

func main() {
	test := []int{1,2,3,4,4,4,4,5,6,7,8}
	fmt.Println("Count 4 in ", test, ": ", count(test, 4))
	test = []int{4,4,4,4}
	fmt.Println("Count 4 in ", test, ": ", count(test, 4))
	test = []int{1,2,3,4,4,4,4,6,7,8}
	fmt.Println("Count 5 in ", test, ": ", count(test, 5))
}
