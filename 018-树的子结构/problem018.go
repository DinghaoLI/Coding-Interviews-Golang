package main

import (
	"fmt"
)

type TreeNode struct {
	Val  int
	Left *TreeNode
	Right *TreeNode
}



func hasSubRootOrTree(p *TreeNode, c *TreeNode) bool {
	if c == nil { return true }
	if p == nil { return false }

	if p.Val == c.Val {
		if hasSub(p, c) {
			return true
		}
	}

	return hasSubRootOrTree(p.Left, c) || hasSubRootOrTree(p.Right, c)
}

func hasSub(p *TreeNode, c *TreeNode) bool {
	if c == nil { return true }
	if p == nil { return false }

	if p.Val != c.Val {
		return true
	}

	return hasSub(p.Left, c.Left) && hasSub(p.Right, c.Right)
}

func main() {
	//test
	root := &TreeNode{1, &TreeNode{2, &TreeNode{4, nil, nil}, &TreeNode{5, nil, nil}}, &TreeNode{3, nil, nil}}
	sub1 := &TreeNode{1, &TreeNode{2, nil, nil}, nil}
	sub2 := &TreeNode{2, &TreeNode{4, nil, nil}, &TreeNode{5, nil, nil}}
	sub3 := &TreeNode{2, nil, nil}

	fmt.Println(hasSubRootOrTree(root, sub1), hasSubRootOrTree(root, sub2), hasSubRootOrTree(root, sub3))



}
