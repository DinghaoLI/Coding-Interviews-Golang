package problem010

import "errors"

func getMostFreq(nums []int) (int, error){
	if len(nums) == 0 {
		return -1, errors.New("Array is empty")
	}
	count := 1
	value := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] == value {
			count++
		} else {
			count--
			if count == 0 {
				value = nums[i]
				count = 1
			}
		}
	}
	return value, nil
}