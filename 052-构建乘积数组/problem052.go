package main

import (
	"fmt"
)

func multiArray(nums []int) []int {
	res := make([]int, len(nums))
	tmp := 1
	for i := range res {
		res[i] = tmp
		tmp *= nums[i]
	}

	tmp = 1
	for i := range res {
		res[len(res)-1-i] *= tmp
		tmp *=  nums[len(res)-1-i]
	}

	return res

}

func main() {

	list := []int{1,2,3,4,5}
	fmt.Println(list, " => ", multiArray(list))


}
