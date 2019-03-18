package problem007

import (
	"fmt"
	"errors"
	"../utils"
)

type Queue struct {
	in utils.Stack
	out utils.Stack
}

func (q *Queue) IsEmpty() bool {
	return q.in.IsEmpty() && q.out.IsEmpty()
}

func (q *Queue) Push(value interface{}) {
	q.in.Push(value)
}

func (q *Queue) Pop() (interface{}, error) {
	if q.IsEmpty() {
		return nil, errors.New("Queue is empty")
	}

	var value interface{}

	if !q.out.IsEmpty() {
		value, _ = q.out.Pop()
		return value, nil
	}

	for !q.in.IsEmpty() {
		value, _ = q.in.Pop()
		q.out.Push(value)
	}
	value, _ = q.out.Pop()
	return value, nil
}


func main() {
	var myQueue Queue

    myQueue.Push(1)
    myQueue.Push(2)
    fmt.Println(myQueue.Pop())
    fmt.Println(myQueue.Pop())
    myQueue.Push(3)
    myQueue.Push(4)
    fmt.Println(myQueue.Pop())
    fmt.Println(myQueue.Pop())
    fmt.Println(myQueue.Pop())
}








