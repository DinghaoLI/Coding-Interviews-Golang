package main

import (
	"fmt"
)

// dp[i] 代表以i结尾时最大连续子数组的最大和
// dp[i] 初始值为nums[0]
// 计算dp[i]时
// 如果dp[i-1]>0: 那我们放心的加进来所以dp[i] = nums[i]+dp[i-1]
// 否则：dp[i]取自己 dp[i] = nums[i]
// 再利用全局变量，保存出出现过的最大值
func MaxSubset(nums []int) int {
	if len(nums) == 0 { return 0 }
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	res := nums[0]
	for i:=1; i<len(nums); i++{
		if dp[i-1] > 0 {
			dp[i] = nums[i]+dp[i-1]
		} else {
			dp[i] = nums[i]
		}

		if dp[i] > res {
			res = dp[i]
		}
	}
	return res
}

// 精简一些
// dp[i] 代表以i结尾时最大连续子数组的最大和（同时）
// 复制nums到dp
// 如果dp[i-1]>0那就把dp[i]=dp[i-1]+dp[i]
// 再利用全局变量，保存出出现过的最大值

func MaxSubset2(nums []int) int {
	if len(nums) == 0 { return 0 }
	dp := make([]int, len(nums))
	copy(dp, nums)
	max := dp[0]
	for i := 1; i < len(dp); i++ {
		if dp[i-1]>0{
			dp[i] = dp[i-1]+dp[i]
		}
		if dp[i]>max {
			max = dp[i]
		}
	}
	return max
}

func max(a,b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	test := []int{1,-2,3,10,-4,7,2,-5}
	fmt.Println(test)
	fmt.Println(MaxSubset(test))
	fmt.Println(MaxSubset2(test))
}