package main

import (
	"fmt"
)

type RandNodeList struct {
	Val  int
	Next *RandNodeList
	Rand *RandNodeList
}

func CopyRandRandNodeList(head *RandNodeList) *RandNodeList {
	if head == nil { return nil }
	m := make(map[*RandNodeList]*RandNodeList)
	cur := head
	for cur != nil {
		if _, ok := m[cur]; !ok {
			m[cur] = &RandNodeList{cur.Val, nil, nil}
		}
		if cur.Next != nil {
			if _, ok := m[cur.Next]; !ok {
				m[cur.Next] = &RandNodeList{cur.Next.Val, nil, nil}
			}
			m[cur].Next = m[cur.Next]
		}
		
		if cur.Rand != nil {
			if _, ok := m[cur.Rand]; !ok {
				m[cur.Rand] = &RandNodeList{cur.Rand.Val, nil, nil}
			}
			m[cur].Rand = m[cur.Rand]
		}
		cur = cur.Next
	}
	return m[head]
}

func print(head *RandNodeList){
	for head!=nil{
		fmt.Printf("%d Ram: %d -> ", head.Val, head.Rand.Val)
		head=head.Next
	}
}

func main() {
	//test
	l3 := &RandNodeList{3, nil, nil}
	l2 := &RandNodeList{2, l3, nil}
	l1 := &RandNodeList{1, l2, nil}
	l1.Rand = l3
	l2.Rand = l2
	l3.Rand = l1
	print(l1)
	fmt.Println("")
	fmt.Println("Copy: ")
	print(CopyRandRandNodeList(l1))
	fmt.Println("")

}











