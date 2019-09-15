package main

import (
	"fmt"
	"math"
)

func myAtoi(str string) int {
	res := 0
	flag := 1
	start := 0
	// space
	for start <= len(str)-1 && str[start] == ' ' {
		start++
	}
	// +/-
	if start <= len(str)-1 {
		if str[start] == '-' {
			flag = -1
			start++
		} else if str[start] == '+' {
			start++
		}
	}
	// is digital ?
	for start <= len(str)-1 && isDigital(str[start]) {
		if res == 0 {
			res += int(str[start] - '0')
			start++
		} else {
			res = res*10 + int((str[start] - '0'))
			start++
		}
		// overflow int32
		if res*flag > math.MaxInt32 {
			return math.MaxInt32
		} else if res*flag < math.MinInt32 {
			return math.MinInt32
		}
	}

	res *= flag
	return res
}

func isDigital(b byte) bool {
	if b <= '9' && b >= '0' {
		return true
	}
	return false
}

func main() {

	fmt.Println("      42", " => ", myAtoi("      42"))
	fmt.Println("      42asdfsdf", " => ", myAtoi("      42asdfsdf"))
	fmt.Println("      -42asdfsdf", " => ", myAtoi("      -42asdfsdf"))
	fmt.Println("9223372036854775808", " => ", myAtoi("9223372036854775808"))
	fmt.Println("-9223372036854775808", " => ", myAtoi("-9223372036854775808"))

}
