package main

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Serialize(root *TreeNode) string {
	if root == nil {
		return "#,"
	}
	str := ""
	str = SerializeTree(root, str)
	return str
}

func SerializeTree(root *TreeNode, str string) string {
	if root == nil {
		str += "#,"
		return str
	}

	str += strconv.Itoa(root.Val) + ","
	str += SerializeTree(root.Left, "")
	str += SerializeTree(root.Right, "")
	return str
}

func Deserialize(str string) *TreeNode {
	if str == "" {
		return nil
	}
	return DeserializeTree(str)
}

func DeserializeTree(str string) *TreeNode {
	index := 0
	var recur func(str string) *TreeNode
	recur = func(str string) *TreeNode {
		if str[index] == '#' {
			index += 2
			return nil
		}
		num := 0
		for str[index] != ',' && index < len(str) {
			num = num*10 + int(str[index] - '0')
			index++
		}
		index++

		root := &TreeNode{num, nil, nil}
		root.Left = recur(str)
		root.Right = recur(str)
		return root
	}
	root := recur(str)
	return root

}

func print(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Printf("%d -> ", root.Val)
	print(root.Left)
	print(root.Right)
}

func main() {
	// q1    1
	// q2  2   3
	// q1 4 5 6 7
	root := &TreeNode{1, &TreeNode{2, &TreeNode{4, nil, nil}, &TreeNode{5, nil, nil}}, &TreeNode{3, &TreeNode{6, nil, nil}, &TreeNode{7, nil, nil}}}

	fmt.Println("Output: ", Serialize(root))
	fmt.Println("")
	fmt.Println("before")
	print(root)
	fmt.Println("")
	root = Deserialize(Serialize(root))
	fmt.Println("after")
	print(root)
	fmt.Println("")

}
