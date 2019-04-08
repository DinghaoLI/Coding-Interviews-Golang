package main

import (
	"fmt"
)

// 递归的思想
// 易错点
// 多指返回 定义 var recur func(root *TreeNode) (int, bool)
// 返回值第一个为深度，第二个为是否平衡
//
// 只要遇到以下情况就返回false
// 1.左右子树只要有一个不平衡
// 2.左右子树深度相差大于一
// 
// 注意返回当前深度时为 max(ldepth, rdepth)+1
func isBalanced(root *TreeNode) bool {
    var recur func(node *TreeNode) (int, bool)
    recur = func(node *TreeNode) (int, bool) {
        if node == nil { return 0, true }
        ldepth, lbalance := recur(node.Left)
        rdepth, rbalance := recur(node.Right)
        if !lbalance || !rbalance || abs(ldepth-rdepth)>1 { 
            return max(ldepth, rdepth)+1, false 
        } 
        return max(ldepth, rdepth)+1, true 
    }
    _, balanced := recur(root)
    return balanced
}

// 精简版
func isBalanced(root *TreeNode) bool {
    if root == nil { return true }
    var recur func(root *TreeNode) (int, bool)
    recur = func(root *TreeNode) (int, bool) {
        if root == nil { return 0, true }
        rightD, rightB := recur(root.Right)
        leftD, leftB := recur(root.Left)
        return max(rightD, leftD)+1, abs(rightD-leftD)<=1 && rightB && leftB
    }
    _, res := recur(root)
    return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}


// 测试地址
// https://leetcode.com/problems/balanced-binary-tree/
