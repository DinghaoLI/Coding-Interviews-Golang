package main

import (
	"fmt"
)



type TreeNode struct {
	Val int
    Left *TreeNode
    Right *TreeNode
}

func TreeToList(root *TreeNode) (*TreeNode, *TreeNode){
	head, tail := root, root
    if root == nil { return head, tail }
    if root.Left == nil && root.Right == nil {
    	head.Left, tail.Right = root, root
        return head, tail
    }
  	if root.Left != nil {
    	leftHead, leftTail := TreeToList(root.Left)
        head = leftHead
        root.Left = leftTail
        leftTail.Right = root  
    }
    if root.Right != nil {
    	rightHead, rightTail := TreeToList(root.Right)
		tail = rightTail
        root.Right = rightHead
        rightHead.Left = root
    }
    head.Left, tail.Right = head, tail
	return head, tail
}



func printRight(root *TreeNode){
	for root.Right != root {
		fmt.Printf("%d ", root.Val)
		root = root.Right
	}
	fmt.Printf("%d ", root.Val)
}

func printLeft(root *TreeNode){
	for root.Left != root {
		fmt.Printf("%d ", root.Val)
		root = root.Left
	}
	fmt.Printf("%d ", root.Val)
}

func main() {
	//test
	root := &TreeNode{5, nil, nil}
	l11 := &TreeNode{3, nil, nil}
	l12 := &TreeNode{8, nil, nil}
	l21 := &TreeNode{1, nil, nil}
	l22 := &TreeNode{4, nil, nil}
	l23 := &TreeNode{6, nil, nil}
	l24 := &TreeNode{9, nil, nil}
	root.Left, root.Right = l11, l12
	l11.Left, l11.Right = l21, l22
	l12.Left, l12.Right = l23, l24
	head, tail := TreeToList(root)
	printLeft(tail)
	fmt.Printf("\n")
	printRight(head)
	fmt.Printf("\n")

	fmt.Printf("\n")

	root = &TreeNode{3, &TreeNode{1, nil, nil}, &TreeNode{4, nil, nil}}
	head, tail = TreeToList(root)
	printLeft(tail)
	fmt.Printf("\n")
	printRight(head)
	fmt.Printf("\n")

}
