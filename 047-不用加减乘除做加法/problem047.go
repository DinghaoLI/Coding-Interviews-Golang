package main

import (
	"fmt"
)

func Add(n, m int) int {
	tmp := 0
	for m != 0 {
		tmp = n ^ m
		m = (n & m) << 1
		n = tmp
	}
	return n
}

func main() {

	fmt.Println("5+7: ", Add(5, 7))
	fmt.Println("50+70: ", Add(50, 70))
	fmt.Println("15+7: ", Add(15, 7))
}
