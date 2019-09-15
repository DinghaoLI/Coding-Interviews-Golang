package main

import (
	"fmt"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
	Parent *TreeNode
}

func getNext(node *TreeNode) *TreeNode {
	if node == nil {
		return node
	}
	//如果节点有右子树，那么它的下一个节点就是它的右子树中最左边的节点
	if node.Right != nil {
		node = node.Right
		for node.Left != nil {
			node = node.Left
		}
		return node
	}
	// 先取目标的父节点
	p := node.Parent
	n := node
	for p != nil {
		// 如果p节点是p的父节点的右节点 =》 继续向上找
		if n == p.Right {
			n = p
			p = p.Parent
			continue
		}
		// p是p父节点的左节点 =》 返回父节点
		return p
	}
	// 目标节点没有下一个节点
	return nil
}

func inOrder(root *TreeNode) {
	if root == nil {
		return
	}
	inOrder(root.Left)
	fmt.Printf("%d -> ", root.Val)
	inOrder(root.Right)
}

func main() {
	n1 := &TreeNode{1, nil, nil, nil}
	n2 := &TreeNode{2, nil, nil, nil}
	n3 := &TreeNode{3, nil, nil, nil}
	n4 := &TreeNode{4, nil, nil, nil}
	n5 := &TreeNode{5, nil, nil, nil}
	n6 := &TreeNode{6, nil, nil, nil}
	n7 := &TreeNode{7, nil, nil, nil}
	n8 := &TreeNode{8, nil, nil, nil}
	n9 := &TreeNode{9, nil, nil, nil}

	n4.Left, n4.Right, n4.Parent = n2, n6, n8
	n2.Left, n2.Right, n2.Parent = n1, n3, n4
	n1.Parent = n2
	n3.Parent = n2
	n6.Left, n6.Right, n6.Parent = n5, n7, n4
	n5.Parent = n6
	n7.Parent = n6
	n8.Left, n8.Right = n4, n9
	n9.Parent = n8

// Example: 
//                8
//          4	        9
//       2	   6
//     1   3  5  7	

	fmt.Println("Inorder: ")
	inOrder(n8)
	fmt.Println("")
	fmt.Println(n1.Val, " getNext : ", getNext(n1).Val)
	fmt.Println(n2.Val, " getNext : ", getNext(n2).Val)
	fmt.Println(n3.Val, " getNext : ", getNext(n3).Val)
	fmt.Println(n4.Val, " getNext : ", getNext(n4).Val)
	fmt.Println(n5.Val, " getNext : ", getNext(n5).Val)
	fmt.Println(n6.Val, " getNext : ", getNext(n6).Val)
	fmt.Println(n7.Val, " getNext : ", getNext(n7).Val)
	fmt.Println(n8.Val, " getNext : ", getNext(n8).Val)
	fmt.Println(n9.Val, " getNext : ", getNext(n9))


}




