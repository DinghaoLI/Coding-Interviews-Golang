package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func print(root *TreeNode) {
	var q1 []*TreeNode
	q1 = append(q1, root)
	for len(q1) != 0 {
		for len(q1) != 0 {
			node := q1[0]
			q1 = q1[1:len(q1)]
			fmt.Printf("%d ", node.Val)
			if node.Left != nil {
				q1 = append(q1, node.Left)
			}
			if node.Right != nil {
				q1 = append(q1, node.Right)
			}
		}
	}
}

func main() {
	// q1    1
	// q2  2   3
	// q1 4 5 6 7
	root := &TreeNode{1, &TreeNode{2, &TreeNode{4, nil, nil}, &TreeNode{5, nil, nil}}, &TreeNode{3, &TreeNode{6, nil, nil}, &TreeNode{7, nil, nil}}}
	print(root)
	fmt.Printf("\n")

	root.Right.Left.Left = &TreeNode{8, nil, nil}
	root.Right.Left.Right = &TreeNode{9, nil, nil}
	root.Right.Right.Left = &TreeNode{10, nil, nil}
	root.Right.Right.Right = &TreeNode{11, nil, nil}
	print(root)
	fmt.Printf("\n")
}
