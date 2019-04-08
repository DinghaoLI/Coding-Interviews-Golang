package main

import (
	"fmt"
)

// 把target-v当作key存入下次遇到target-v就说明找到了
func twoSum(nums []int, target int) []int {
    m := make(map[int]int)
    for i, v := range nums {
        if val, ok := m[v]; ok {
            return []int{val, i}
        }
        m[target-v] = i
    }
    return []int{}
}
// 测试地址
// https://leetcode.com/problems/two-sum/


