package main

import (
	"fmt"
)

type Solution struct {
	str   string
	count []int
}

func (s *Solution) Insert(char byte) {
	s.str += string(char)
	s.count[char]++
}

func (s *Solution) GetOnceAppear() byte {
	for i, v := range s.count {
		if v == 1 {
			return byte(i)
		}
	}
	return '#'
}

func main() {
	sol := &Solution{"", make([]int, 256)}

	sol.Insert('g')
	fmt.Println(sol.str, " GetOnceAppear: ", string(sol.GetOnceAppear()))
	sol.Insert('o')
	fmt.Println(sol.str, " GetOnceAppear: ", string(sol.GetOnceAppear()))
	sol.Insert('o')
	fmt.Println(sol.str, " GetOnceAppear: ", string(sol.GetOnceAppear()))
	sol.Insert('g')
	fmt.Println(sol.str, " GetOnceAppear: ", string(sol.GetOnceAppear()))
	sol.Insert('l')
	fmt.Println(sol.str, " GetOnceAppear: ", string(sol.GetOnceAppear()))
	sol.Insert('e')
	fmt.Println(sol.str, " GetOnceAppear: ", string(sol.GetOnceAppear()))

}
