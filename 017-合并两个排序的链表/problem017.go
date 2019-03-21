package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func print(head *ListNode) {
	for head!=nil{
		fmt.Printf("%d -> ", head.Val)
		head=head.Next
	}
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{}
	cur := res
	for l1 != nil || l2 != nil{
		if l1 == nil {
			cur.Next = l2
			break
		} else if l2 == nil {
			cur.Next = l1
			break
		}
		if l1.Val > l2.Val {
			cur.Next = l2
			cur = cur.Next
			l2 = l2.Next
		} else {
			cur.Next = l1
			cur = cur.Next
			l1 = l1.Next
		}
	}
	return res.Next
}

func main() {
	//test
	l3 := &ListNode{6, nil}
	l2 := &ListNode{4, l3}
	l1 := &ListNode{2, l2}
	
	print(l1)

	fmt.Println("\n")
	
	n3 := &ListNode{5, nil}
	n2 := &ListNode{3, n3}
	n1 := &ListNode{1, n2}
	
	print(n1)

	fmt.Println("\n")
	print(mergeTwoLists(l1, n1))

	fmt.Println("\n")
}
