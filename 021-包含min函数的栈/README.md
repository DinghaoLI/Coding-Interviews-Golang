# 题意
题目描述

定义栈的数据结构，请在该类型中实现一个能够得到栈最小元素的min函数。

# 分析
思路很简单，我们维持两个栈，

数据栈data，存储栈的数据用于常规的栈操作

最小栈min，保存每次push和pop时候的最小值，

在push-data栈的时候，将当前最小数据压入，

在pop-data栈的时候，将min栈栈顶的最小数据弹出

这样保证min栈中存储着当前现场的最小值，并随着数据栈的更新而更新

## 参考Leetcode原题

[LeetCode 155. Min Stack](https://leetcode.com/problems/min-stack/)

```golang
// 易错点
// 对于每一个的item的min，如果，前一个元素的min比自己的x还小，那么自己的min就储存之前的元素min，反之，元素min等于x

// MinStack 是可以返回最小值的栈
type MinStack struct {
	stack []item
}
type item struct {
	min, x int
}

// Constructor 构造 MinStack
func Constructor() MinStack {
	return MinStack{}
}

// Push 存入数据
func (s *MinStack) Push(x int) {
	min := x
	if len(s.stack) > 0 && s.GetMin() < x {
		min = s.GetMin()
	}
	s.stack = append(s.stack, item{min: min, x: x})
}

// Pop 抛弃最后一个入栈的值
func (s *MinStack) Pop() {
	s.stack = s.stack[:len(s.stack)-1]
}

// Top 返回最大值
func (s *MinStack) Top() int {
	return s.stack[len(s.stack)-1].x
}

// GetMin 返回最小值
func (s *MinStack) GetMin() int {
	return s.stack[len(s.stack)-1].min
}

```