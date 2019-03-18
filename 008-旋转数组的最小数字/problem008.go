package problem008

func minNumberInRotateArray(array []int) int {
	low, high := 0, len(array)-1
	if array[low] < array[high] {
		return array[low]
	}
	mid := (low + high) / 2
	for array[low] >= array[high] {
		// array[low] >= array[high] 所以此时high为最小值
		if high-low == 1 {
			mid = high
			break
		}
		mid = (low + high) / 2

		// array[low] array[mid] array[high]三者相等
		// 无法确定中间元素是属于前面还是后面的递增子数组
		// 只能顺序查找
		if array[low] == array[mid] && array[mid] == array[high] {
			return MinOrder(array, low, high)
		}


		if array[mid] >= array[low] {
			low = mid
		} else {
			high = mid
		}
	}
	return array[mid]
}

// 顺序查找
func MinOrder(array []int, low, high int) int {
	min := array[low]
	for i := low; i <= high; i++ {
		if array[i] < min {
			min = array[i]
		}
	}
	return min
}
