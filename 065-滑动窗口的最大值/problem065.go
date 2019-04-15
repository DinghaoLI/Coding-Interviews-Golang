// deque 双向队列
func maxSlidingWindow(nums []int, k int) []int {
	var res []int
	var window []int
	for index := 0; index < len(nums); index++ {
        // 最大值已经超出窗口范围
		if index >= k && window[0] <= index - k {
			window = window[1:]
		}
        // 当前值比队列末尾大则替换末尾
		for len(window) > 0 && nums[window[len(window) - 1]] < nums[index] {
			window = window[:len(window) - 1]
		}
        // 添加元素进队列
		window = append(window, index)
        // 记录当前最大值
		if index >= k - 1 {
			res = append(res, nums[window[0]])    
		}
	}
	return res
}

// 测试地址
// https://leetcode.com/problems/sliding-window-maximum/