package main

import (
	"fmt"
)

type TreeNode struct {
	Val  int
	Left *TreeNode
	Right *TreeNode
}



func MirrorTree(p *TreeNode) {
	if p == nil { return }
	p.Left, p.Right = p.Right, p.Left
	MirrorTree(p.Left)
	MirrorTree(p.Right)
}

func Print(root *TreeNode){
	if root == nil { return }
	fmt.Printf("%d ", root.Val)
	Print(root.Left)
	Print(root.Right)
}

func main() {
	//test
	root := &TreeNode{1, &TreeNode{2, &TreeNode{4, nil, nil}, &TreeNode{5, nil, nil}}, &TreeNode{3, nil, nil}}
	Print(root)
	fmt.Printf("\n")
	MirrorTree(root)
	Print(root)
	fmt.Printf("\n")

	root = &TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}}
	Print(root)
	fmt.Printf("\n")
	MirrorTree(root)
	Print(root)
	fmt.Printf("\n")

}
