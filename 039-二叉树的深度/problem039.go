package main

import (
	"fmt"
)

type TreeNode {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
    max := 0
	var dfs func(root *TreeNode, level int)
	dfs = func(root *TreeNode, level int){
		if root == nil {
			return 
		}
		if level > max {
			max = level
		}
		dfs(root.Left, level+1)
		dfs(root.Right, level+1)
	}
	dfs(root, 1)
	return max
}

// 测试地址
// https://leetcode.com/problems/maximum-depth-of-binary-tree/

