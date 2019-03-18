package problem007

import (
    "testing"
)

func TestQueue_IsEmpty(t *testing.T) {
    var myQueue Queue
    myQueue.Push(1)
    myQueue.Push("test")
    if myQueue.IsEmpty() == false {
        t.Log("Pass Queue.IsEmpty")
    } else {
        t.Error("Failed Queue.IsEmpty")
    }
}

func TestQueue_Push(t *testing.T) {
    var mQueue Queue
    mQueue.Push(3)
    mQueue.Push(4)
    mQueue.Push(5)
    mQueue.Push(6)
    if mQueue.IsEmpty() == false {
        t.Log("Pass Queue.Push")
    } else {
        t.Error("Failed Queue.Push")
    }
}

func TestQueue_Pop(t *testing.T) {
    var mQueue Queue
    if _, err := mQueue.Pop(); err == nil {
        t.Error("Failed Queue.Pop")
    }

    mQueue.Push(3)
    mQueue.Push(4)


    if value, _ := mQueue.Pop(); value == 3 {
        t.Log("Pass Queue.Pop")
    } else {
        t.Errorf("Failed Queue.Pop, value should be %d", 3)
    }
    if value, _ := mQueue.Pop(); value == 4 {
        t.Log("Pass Queue.Pop")
    } else {
        t.Errorf("Failed Queue.Pop, value should be %d", 4)
    }

    if _, err := mQueue.Pop(); err == nil {
        t.Error("Failed Queue.Pop")
    }

    mQueue.Push(5)

    if value, _ := mQueue.Pop(); value == 5 {
        t.Log("Pass Queue.Pop")
    } else {
        t.Errorf("Failed Queue.Pop, value should be %d", 5)
    }

    if _, err := mQueue.Pop(); err == nil {
        t.Error("Failed Queue.Pop")
    }


}