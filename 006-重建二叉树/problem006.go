package main

import (
	"fmt"
)

type TreeNode struct{
	Val int
	Left *TreeNode
	Right *TreeNode
}

func printPreOrder(root *TreeNode){
	if root != nil {
		fmt.Printf("%d ", root.Val)
		printPreOrder(root.Left)
		printPreOrder(root.Right)
	}
}

func printInOrder(root *TreeNode){
	if root != nil {
		printInOrder(root.Left)
		fmt.Printf("%d ", root.Val)
		printInOrder(root.Right)
	}
}

func reConstructBinaryTree(pre []int, in []int) *TreeNode {
	if len(pre) != len(in) || len(pre) == 0 { 
		return nil 
	}
	// find root and root Index in inOrder
	rootVal := pre[0]
	rootIndex := 0
	for i := 0; i < len(in); i++ {
		if in[i] == rootVal {
			rootIndex = i
		}
	}
	// pre and in for left and right 
	inL, inR := in[:rootIndex], in[rootIndex+1:]
	preL, preR := pre[1:rootIndex+1], pre[rootIndex+1:]
	// revursive
	left := reConstructBinaryTree(preL, inL)
	right := reConstructBinaryTree(preR, inR)
	return &TreeNode{Val: rootVal, Left: left, Right: right}
}


func main() {
	// example
	pre := []int{1,2,4,7,3,5,6,8}
	in := []int{4,7,2,1,5,3,6,8}

	fmt.Println("preOder: ", pre)
	fmt.Println("inOrder: ", in)

	// Reconstruct
	fmt.Println("\nReconstruct Binary Tree... \n ",)
	root := reConstructBinaryTree(pre, in)

	// test 
	fmt.Printf("preOder from Tree reconstructed:  ")
	printPreOrder(root)
	fmt.Printf("\n")

	fmt.Printf("inOder from Tree reconstructed:  ")
	printInOrder(root)
	fmt.Printf("\n")
}
























