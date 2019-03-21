package main

import (
	"fmt"
	"../utils"
)


func stackOrder(in []int, out []int) bool {
	var s utils.Stack
	if len(out) != len(in) { return false }
	pi, po := 0, 0
	for pi<len(in) {
		if out[po] != in[pi] {
			s.Push(in[pi])
			pi++
			continue
		}
		po++
		pi++
		value, _ := s.Top()
		for po<len(out) && out[po] == value {
			s.Pop()
			po++
			value, _ = s.Top()
		} 
	}
	return po == len(out)
}

func main() {
	in := []int{1,2,3,4,5}
	fmt.Println("stack in : ", in)

	out1 := []int{4,5,3,2,1}
	out2 := []int{4,3,5,1,2}
	out3 := []int{5,4,3,2,1}
	fmt.Println("out1: ", out1, "  ", stackOrder(in, out1))
	fmt.Println("out2: ", out2, "  ", stackOrder(in, out2))
	fmt.Println("out3: ", out3, "  ", stackOrder(in, out3))		

}
