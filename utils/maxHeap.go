package utils

import (
	"fmt"
	"math"
)

type MaxHeap struct {
	Element []int
}



// MaxHeap constructor
func NewMaxHeap() *MaxHeap {
	first := math.MaxInt64
	h := &MaxHeap{Element: []int{first}}
	return h
}

// Length of Maxheap
func (H *MaxHeap) Length() int {
	return len(H.Element) - 1
}

// Get the Maximum of the Maxheap
func (H *MaxHeap) Max() (int, error) {
	if len(H.Element) > 1 {
		return H.Element[1], nil
	}
	return -1, fmt.Errorf("heap is empty")
}

// Get the Maximum of the Maxheap
func (H *MaxHeap) GetMax() int {
	if H.Length() > 0 {
		return H.Element[1]
	}
	return math.MaxInt64
}

// Inserting items requires ensuring the nature of the Maxheap
func (H *MaxHeap) Insert(v int) {
	H.Element = append(H.Element, v)
	i := len(H.Element) - 1
	for ; (H.Element[i/2]) < v; i /= 2 {
		H.Element[i] = H.Element[i/2]
	}

	H.Element[i] = v
}

// Delete and return the Maximum
func (H *MaxHeap) DeleteMax() (int, error) {
	if len(H.Element) <= 1 {
		return -1, fmt.Errorf("MaxHeap is empty")
	}
	MaxElement := H.Element[1]
	lastElement := H.Element[len(H.Element)-1]
	var i, child int
	for i = 1; i*2 < len(H.Element); i = child {
		child = i * 2
		if child < len(H.Element)-1 && H.Element[child+1] > H.Element[child] {
			child++
		}
		if lastElement < H.Element[child] {
			H.Element[i] = H.Element[child]
		} else {
			break
		}
	}
	H.Element[i] = lastElement
	H.Element = H.Element[:len(H.Element)-1]
	return MaxElement, nil
}
