package main

import (
	"fmt"
	"strconv"
)

func getMinStr(nums []int) string {
	quickSort(nums, 0, len(nums)-1)
	noZero := 0
	for nums[noZero] == 0 {
		noZero++
	}
	res := ""
	fmt.Println(nums)
	for i := noZero; i <= len(nums)-1; i++ {
		res += strconv.Itoa(nums[i])
	}
	return res

}

// 要点：
// 注意sup的条件，确定一个数应该在前面还是后面
func quickSort(nums []int, left int, right int) {
	if left < right {
		tmp := nums[left]
		l, r := left, right
		for {
			// 先从右向左!!!
			for l < r && sup(nums[r],tmp) {
				r--
			}
			for l < r && inf(nums[l],tmp) {
				l++
			}
			if l >= r {
				break
			}
			nums[l], nums[r] = nums[r], nums[l]
		}
		nums[left], nums[l] = nums[l], nums[left]
		quickSort(nums, left, l-1)
		quickSort(nums, l+1, right)
	}
}

func sup(a, b int) bool {
	aStr := strconv.Itoa(a)
	bStr := strconv.Itoa(b)

	if aStr+bStr >= bStr+aStr {
		return true
	}
	return false
}

func inf(a, b int) bool {
	aStr := strconv.Itoa(a)
	bStr := strconv.Itoa(b)
	if aStr+bStr <= bStr+aStr {
		return true
	}
	return false
}

type bytes [][]byte

func (b bytes) Less(i, j int) bool {
	size := len(b[i]) + len(b[j])

	bij := make([]byte, 0, size)
	bij = append(bij, b[i]...)
	bij = append(bij, b[j]...)

	bji := make([]byte, 0, size)
	bji = append(bji, b[j]...)
	bji = append(bji, b[i]...)

	for k := 0; k < size; k++ {
		if bij[k] > bji[k] {
			return true
		} else if bij[k] < bji[k] {
			return false
		}
	}

	return false
}


func main() {
	list := []int{10, 9, 9, 111, 111, 222,9, 7, 7, 6, 5, 5, 4, 4, 3, 2, 1, 0, 1, 2, 3, 4, 5}
	fmt.Println(getMinStr(list))
	list = []int{3,30,34,5,9}
	fmt.Println(getMinStr(list))
}
















