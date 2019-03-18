package main

import (
	"fmt"
)

type NodeList struct{
	Val int
	Next *NodeList
}

// just print it form tail to head and do not modify the original NodeList

func printListFromTailToHead(head *NodeList) {
	if head != nil {
		printListFromTailToHead(head.Next)
		fmt.Printf("%d -> ", head.Val)
	}
}

func main() {
	// example
	n3 := &NodeList{3, nil}
	n2 := &NodeList{2, n3}	
	n1 := &NodeList{1, n2}

	fmt.Printf("\n NodeList 1 -> 2 -> 3 \n")

	// test
	fmt.Printf("\n Output: ")
	printListFromTailToHead(n1)	
	fmt.Printf("\n \n")


}
