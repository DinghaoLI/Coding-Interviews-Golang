package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getKth(root *TreeNode, k int) *TreeNode {
	if root == nil || k<1 {
		return nil
	}
	var res *TreeNode
	found := false
	var inOrder func(*TreeNode)
	inOrder = func(node *TreeNode) {
		if found == true || node == nil {
			return
		}
		inOrder(node.Left)
		if k == 1 {
			res = node
			found = true
		}
		k--
		inOrder(node.Right)
	}
	inOrder(root)	
	return res
}


func print(root *TreeNode) {
	if root == nil {
		return
	}
	print(root.Left)
	fmt.Printf("%d -> ", root.Val)
	print(root.Right)
}

func main() {
	// q1    5
	// q2  3   7
	// q1 1 4 6 8
	root := &TreeNode{5, &TreeNode{3, &TreeNode{1, nil, nil}, &TreeNode{4, nil, nil}}, &TreeNode{7, &TreeNode{6, nil, nil}, &TreeNode{8, nil, nil}}}
	print(root)
	fmt.Println("")
	fmt.Println("Get 1th: ", getKth(root, 1).Val)
	fmt.Println("Get 2th: ", getKth(root, 2).Val)
	fmt.Println("Get 3th: ", getKth(root, 3).Val)
	fmt.Println("Get 4th: ", getKth(root, 4).Val)
	fmt.Println("Get 5th: ", getKth(root, 5).Val)
	fmt.Println("Get 6th: ", getKth(root, 6).Val)
	fmt.Println("Get 7th: ", getKth(root, 7).Val)
	fmt.Println("Get 8th: ", getKth(root, 8))


}
