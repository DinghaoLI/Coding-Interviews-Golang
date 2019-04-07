package main

import (
	"fmt"
)

type LinkNode struct {
	Val int
	Next *LinkNode
}

// 长链表先走，实现右对齐
func firstCommon(h1, h2 *LinkNode) *LinkNode {
	start, l1 := h1, 0 
	for start != nil {
		start = start.Next
		l1++
	}
	start, l2 := h2, 0
	for start != nil {
		start = start.Next
		l2++
	}

	s1, s2 := h1, h2
	if l1 > l2 {
		diff := l1-l2
		for s1 != nil && diff >0 {
			s1 = s1.Next
			diff--
		}
	} else if l1 < l2 {
		diff := l2-l1
		for s2 != nil && diff >0 {
			s2 = s2.Next
			diff--
		}
	}
	for s1!=nil && s2!=nil && s1!=s2 {
		s1 = s1.Next
		s2 = s2.Next
	}
	return s1
}

// Hashmap
func firstCommonMap(h1, h2 *LinkNode) *LinkNode {
	m := make(map[*LinkNode]bool)
	for h1 != nil {
		m[h1] = true
		h1 = h1.Next
	}

	for h2 != nil {
		if _, ok := m[h2]; ok {
			return h2
		}
		h2 = h2.Next
	}
	return h2
}

func main() {
	comm := &LinkNode{666, &LinkNode{1, &LinkNode{2, nil}}}
	h1 := &LinkNode{-2, &LinkNode{-1, comm}}
	h2 := &LinkNode{-4, &LinkNode{-3, &LinkNode{-2, comm}}}
	h3 := &LinkNode{-6, comm}

	fmt.Println("Comm: ", (firstCommonMap(h1,h2)))
	fmt.Println("Comm: ", (firstCommonMap(h3,h2)))
	fmt.Println("Comm: ", (firstCommonMap(h3,h1)))

	fmt.Println("Comm: ", (firstCommon(h1,h2)))
	fmt.Println("Comm: ", (firstCommon(h3,h2)))
	fmt.Println("Comm: ", (firstCommon(h3,h1)))
}









