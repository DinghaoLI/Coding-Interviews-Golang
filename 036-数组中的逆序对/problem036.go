package main

import (
	"fmt"
)

func InversePairs(nums []int) int {
	// 用来存储排序好的数组
	tmp := make([]int, len(nums))

	// 合并两个排序好的array
	var merge func (start int, mid int, end int) int
	merge = func (start int, mid int, end int) int {
		if start >= end { return 0 } 
		p1, p2 := mid, end
		k, count := 0, 0
		for p1 >= start && p2 >= mid+1 {
			if nums[p1] <= nums[p2] {
				tmp[k] = nums[p2]
				p2--
				k++
			} else {
				// nums[p1] > nums[p2]
				// 因为两个数组是排好序的，所以说明对于num[p1]来说
				// 至少有p2-mid个逆序对
				tmp[k] = nums[p1]
				count += p2 - mid
				p1--
				k++
			}
		}
		for p1 >= start {
			tmp[k] = nums[p1]
			k++
			p1--
		}
		for p2 >= mid+1 {
			tmp[k] = nums[p2]
			k++
			p2--
		}
		for i := 0; i <= k-1; i++ {
			nums[end-i] = tmp[i] 
		}

		return count
	}
	var sort func (start, end int) int
	sort = func (start, end int) int {
		count := 0
		if start < end {
			mid := (start+end)/2
			count += sort(start, mid)
			count += sort(mid+1, end)
			count += merge(start, mid, end)
			return count
		}
		return 0
	}

	return sort(0, len(nums)-1)
}


func main() {
	test := []int{7,5,6,4}
	fmt.Println(test)
	fmt.Println(" InversePairs: ", InversePairs(test))

	test = []int{7,6,5,4}
	fmt.Println(test)
	fmt.Println(" InversePairs: ", InversePairs(test))
}
















