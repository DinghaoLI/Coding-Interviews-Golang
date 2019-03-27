package main

import (
	"fmt"
)

type NodeList struct {
	Val  int
	Next *NodeList
}

func reverse(head *NodeList) *NodeList {
	if head == nil || head.Next == nil { return head }
	vide := &NodeList{-1, nil}

	for head != nil {
		next := head.Next
		head.Next = vide.Next
		vide.Next = head
		head = next
	}

	return vide.Next

}

func print(head *NodeList){
	for head!=nil{
		fmt.Printf("%d -> ", head.Val)
		head=head.Next
	}
}

func main() {
	//test
	l3 := &NodeList{3, nil}
	l2 := &NodeList{2, l3}
	l1 := &NodeList{1, l2}
	fmt.Println("1 -> 2 -> 3")
	print(reverse(l1))
	fmt.Println("\n")


	l2 = &NodeList{2, nil}
	l1 = &NodeList{1, l2}
	fmt.Println("1 -> 2 ->")
	print(reverse(l1))
	fmt.Println("\n")


	l1 = &NodeList{1, nil}
	fmt.Println("1 -> ")
	print(reverse(l1))
	fmt.Println("\n")

}
