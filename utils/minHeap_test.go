package utils

import (
	"testing"
	"fmt"
)

func Test_MinHeap(t *testing.T) {
	heap := NewMinHeap()
	if heap.Length() == 0 {
		t.Log("Pass")
	} else {
		t.Error("Failed")
	}

	heap.Insert(1)
	heap.Insert(2)
	heap.Insert(3)
	heap.Insert(4)
	heap.Insert(5)


	if heap.Length() == 5 {
		t.Log("Pass")
	} else {
		t.Error("Failed")
	}

	v, _ := heap.Min()
	if v == 1 {
		t.Log("Pass")
	} else {
		t.Error("Failed")
	}

	v, _ = heap.DeleteMin()
	if v == 1 {
		t.Log("Pass")
	} else {
		t.Error("Failed")
	}

	v, _ = heap.DeleteMin()
	if v == 2 {
		t.Log("Pass")
	} else {
		fmt.Println(v)
		t.Error("Failed")
	}
	v, _ = heap.DeleteMin()
	if v == 3 {
		t.Log("Pass")
	} else {
		fmt.Println(v)
		t.Error("Failed")
	}
	v, _ = heap.DeleteMin()
	if v == 4 {
		t.Log("Pass")
	} else {
		fmt.Println(v)
		t.Error("Failed")
	}

}
