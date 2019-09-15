package main

import (
	"fmt"
	"errors"
	. "../utils"
)

func minK(nums []int, k int) ([]int, error){
	res := []int{}
	if k > len(nums) {
		return res, errors.New("k > length of nums")
	}

	maxHeap := NewMaxHeap()

	for _, v := range nums {
		if maxHeap.Length() < k {
			maxHeap.Insert(v)
		} else {
			max, _ := maxHeap.Max()
			if max > v {
				maxHeap.DeleteMax()
				maxHeap.Insert(v)
			}
		}
	}
	for maxHeap.Length()>0 {
		v, _ := maxHeap.DeleteMax()
		res = append(res, v)
	}
	return res, nil
}

func main() {
	test := []int{1,2,3,4,5,6,7,8,9}
	fmt.Println(test)
	fmt.Println(minK(test, 1))
	fmt.Println(minK(test, 2))
	fmt.Println(minK(test, 3))
	fmt.Println(minK(test, 4))
	fmt.Println(minK(test, 5))
}