# 题意

输入一个矩阵，按照从外向里以顺时针的顺序依次打印出每一个数字，

例如， 如果输入如下矩阵：

1 2 3 4 
5 6 7 8 
9 10 11 12 
13 14 15 16

则依次打印出数字

1,2,3,4,8,12,16,15,14,13,9,5,6,7,11,10

## 参考Leetcode原题

[LeetCode 54. Spiral Matrix](https://leetcode.com/problems/spiral-matrix/)

```golang
// 根据题目要求写出规律，我们把螺旋遍历当作扒一层一层皮
// 我们可以设置一个变量layer，用来记录当前在记录第几层。
// 然后用四个变量来记录四个边界，按照同样的规律扒皮，最后扒完算是结束
// 用元素个数保证正确性
func spiralOrder(matrix [][]int) []int {
	var res []int
	n := len(matrix)
	if n==0 {return res}
	m := len(matrix[0])
    // 计算元素个数
	nb := n * m
    // 初始为0层皮
	layer := 0
    // 初始边界
	startN, endN, startM, endM := 0, 0, 0, 0
    // 每放入一个元素，nb就减一，若nb==0，则说明扒皮结束
	for nb > 0 {
        // 按照layer数计算边界
		startN, endN = layer, n-layer-1
		startM, endM = layer, m-layer-1
        // 4个for来扒皮
		for i := startM; i <= endM && nb > 0; i++ {
			res = append(res, matrix[startN][i])
			nb--
		}
		for i := startN + 1; i <= endN && nb > 0; i++ {
			res = append(res, matrix[i][endM])
			nb--
		}
		for i := endM - 1; i >= startM && nb > 0; i-- {
			res = append(res, matrix[endN][i])
			nb--
		}
		for i := endN - 1; i >= startN+1 && nb > 0; i-- {
			res = append(res, matrix[i][startM])
			nb--
		}
        // 层数递增
		layer++
	}
	return res
}
//