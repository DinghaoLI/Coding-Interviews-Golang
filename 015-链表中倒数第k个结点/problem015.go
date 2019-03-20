package main

import (
	"fmt"
)

type NodeList struct {
	Val  int
	Next *NodeList
}

func kthNode(head *NodeList, k int) *NodeList {
	if head == nil {
		return head
	}
	tail := head
	for k > 1 && tail != nil {
		tail = tail.Next
		k--
	}
	if tail == nil {
		return nil
	}
	for tail.Next != nil {
		tail = tail.Next
		head = head.Next
	}
	return head
}

func main() {
	//test
	l3 := &NodeList{3, nil}
	l2 := &NodeList{2, l3}
	l1 := &NodeList{1, l2}
	fmt.Println("1 -> 2 -> 3")

	fmt.Println("reverse 1th")
	fmt.Println(kthNode(l1, 1))
	fmt.Println("reverse 2th")
	fmt.Println(kthNode(l1, 2))
	fmt.Println("reverse 3th")
	fmt.Println(kthNode(l1, 3))
	fmt.Println("reverse 4th")
	fmt.Println(kthNode(l1, 4))

}
