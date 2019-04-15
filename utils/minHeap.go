package utils

import (
	"fmt"
	"math"
)

type MinHeap struct {
	Element []int
}



// MinHeap constructor
func NewMinHeap() *MinHeap {
	first := math.MinInt64
	h := &MinHeap{Element: []int{first}}
	return h
}

// Length of Minheap
func (H *MinHeap) Length() int {
	return len(H.Element) - 1
}

// Get the Minimum of the Minheap
func (H *MinHeap) Min() (int, error) {
	if len(H.Element) > 1 {
		return H.Element[1], nil
	}
	return -1, fmt.Errorf("heap is empty")
}

// Get the Minimum of the Minheap
func (H *MinHeap) GetMin() int {
	if H.Length() > 0 {
		return H.Element[1]
	}
	return math.MinInt64
}

// Inserting items requires ensuring the nature of the Minheap
func (H *MinHeap) Insert(v int) {
	H.Element = append(H.Element, v)
	i := len(H.Element) - 1
	for ; (H.Element[i/2]) > v; i /= 2 {
		H.Element[i] = H.Element[i/2]
	}

	H.Element[i] = v
}

// Delete and return the Minimum
func (H *MinHeap) DeleteMin() (int, error) {
	if len(H.Element) <= 1 {
		return -1, fmt.Errorf("MinHeap is empty")
	}
	MinElement := H.Element[1]
	lastElement := H.Element[len(H.Element)-1]
	var i, child int
	for i = 1; i*2 < len(H.Element); i = child {
		child = i * 2
		if child < len(H.Element)-1 && H.Element[child+1] < H.Element[child] {
			child++
		}
		if lastElement > H.Element[child] {
			H.Element[i] = H.Element[child]
		} else {
			break
		}
	}
	H.Element[i] = lastElement
	H.Element = H.Element[:len(H.Element)-1]
	return MinElement, nil
}
