package main

import (
	"fmt"
)

func isNumber(str string) bool {
	if len(str) == 0 {
		return false
	}
	start := 0
	if str[start] == '+' || str[start] == '-' {
		start++
	}
	if start >= len(str) {
		return false
	}

	numberic := true

	start = scanDigits(str, start)


	if start < len(str) {
		if str[start] == '.' {
			start++
			start = scanDigits(str, start)
			if start == len(str) {
				return true
			}
			if str[start] == 'e' || str[start] == 'E' {
				numberic, start = isE(str, start)

			}
		} else if str[start] == 'e' || str[start] == 'E' {

			numberic, start = isE(str, start)

		} else {
			numberic = false
		}
	}

	return numberic && start == len(str)

}

func scanDigits(str string, start int) int {
	i := start
	for i < len(str) && str[i] <= '9' && str[i] >= '0' {
		i++
	}
	return i
}

func isE(str string, start int) (bool, int) {
	if str[start] != 'e' && str[start] != 'E' {
		return false, start
	}
	start++
	if start >= len(str) {
		return false, start
	}

	if str[start] == '+' || str[start] == '-' {
		start++
	}

	if start >= len(str) {
		return false, start
	}

	start = scanDigits(str, start)

	if start == len(str) {
		return true, start
	}
	return false, start
}

func main() {
	fmt.Println("Should be true")

	test := "+100"
	fmt.Println(test, " is Number ", isNumber(test))
	test = "5e2"
	fmt.Println(test, " is Number ", isNumber(test))
	test = "-123"
	fmt.Println(test, " is Number ", isNumber(test))

	test = "-3.1416"
	fmt.Println(test, " is Number ", isNumber(test))

	test = "-3E-16"
	fmt.Println(test, " is Number ", isNumber(test))

	fmt.Println("")
    fmt.Println("Should be false")

	test = "12e"
	fmt.Println(test, " is Number ", isNumber(test))

	test = "1a3.14"
	fmt.Println(test, " is Number ", isNumber(test))

	test = "1.2.3"
	fmt.Println(test, " is Number ", isNumber(test))

	test = "+-5"
	fmt.Println(test, " is Number ", isNumber(test))

	test = "12e+4.3"
	fmt.Println(test, " is Number ", isNumber(test))

}
