package main

import (
	"../utils"
	"fmt"
)

type Solution struct {
	max *utils.MaxHeap
	min *utils.MinHeap
}

func (s *Solution) Insert(num int) {
	if ((s.max.Length() + s.min.Length()) & 1) == 0 {
		// 偶数数据的情况下
		// 直接将新的数据插入到数组的后半段
		// 即在最小堆中插入元素
		// 此时最小堆中多出一个元素,
		// 即最小元素, 将其弹出后, 压入前半段(即最大堆中)
		if s.max.Length() > 0 && num < s.max.GetMax() {
			s.max.Insert(num)
			val, _ := s.max.DeleteMax()
			s.min.Insert(val)
		} else {
			s.min.Insert(num)
		}
	} else {
		// 奇数情况
		if s.max.Length() > 0 && num < s.max.GetMax() {
			s.max.Insert(num)
		} else {
			s.min.Insert(num)
			val, _ := s.min.DeleteMin()
			s.max.Insert(val)
		}
	}
}

func (s *Solution) GetMedian() float64 {
	size := s.max.Length() + s.min.Length()
	if size == 0 {
		return -1
	}
	var median float64
	if size&1 != 0 {
		median = float64(s.min.GetMin())
	} else {
		median = float64(s.min.GetMin() + s.max.GetMax()) / 2
	}
	return median
}


func main() {

	sol := &Solution{utils.NewMaxHeap(), utils.NewMinHeap()}
	sol.Insert(1)
	fmt.Println("Insert 1")
	fmt.Println(sol.GetMedian())

	sol.Insert(2)
	fmt.Println("Insert 2")
	fmt.Println(sol.GetMedian())

	sol.Insert(3)
	fmt.Println("Insert 3")
	fmt.Println(sol.GetMedian())

	sol.Insert(4)
	fmt.Println("Insert 4")
	fmt.Println(sol.GetMedian())

	sol.Insert(5)
	fmt.Println("Insert 5")
	fmt.Println(sol.GetMedian())

}
