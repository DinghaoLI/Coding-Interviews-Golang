package main

import (
	"fmt"
)

// f(n) = n1*f(10bit-1) + f(n – n1*10bit) + LEFT;
// 其中
// if(n1 == 1)
//     LEFT = n - 10bit+ 1;
// else
//     LEFT = 10bit;

func Ones(n int) int {
	if n == 0 {
		return 0
	}
	if n > 1 && n < 10 {
		return 1
	}
	count := 0
	highest := n
	bit := 0
	for highest >= 10 {
		highest /= 10
		bit++ 
	}

	weight := pow(10, bit)
	if highest == 1 {
		count = Ones(weight - 1) + Ones(n - weight) + ( n - weight + 1)
	} else {
		count = Ones(weight - 1) + Ones(n - highest*weight) + weight
	}

	return count

}

func pow(a,b int) int {
	res := 1
	for i := b; i >0 ; i-- {
		res *= a
	}
	return res
}

func main() {
	
	fmt.Println("10 has ", Ones(10))
	fmt.Println("15 has ", Ones(15))
	fmt.Println("20 has ", Ones(20))
	fmt.Println("100 has ", Ones(100))
	fmt.Println("120 has ", Ones(120))
}





