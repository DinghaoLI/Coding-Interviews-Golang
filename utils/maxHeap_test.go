package utils

import (
	"testing"
	"fmt"
)

func Test_MaxHeap(t *testing.T) {
	heap := NewMaxHeap()
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

	v, _ := heap.Max()
	if v == 5 {
		t.Log("Pass")
	} else {
		t.Error("Failed")
	}

	v, _ = heap.DeleteMax()
	if v == 5 {
		t.Log("Pass")
	} else {
		t.Error("Failed")
	}

	v, _ = heap.DeleteMax()
	if v == 4 {
		t.Log("Pass")
	} else {
		fmt.Println(v)
		t.Error("Failed")
	}

}
