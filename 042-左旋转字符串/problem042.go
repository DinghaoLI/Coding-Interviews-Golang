package main

import (
	"fmt"
)

func rotateString(str string, shift int) string {
	shift = shift%len(str)
	list := []byte(str)
	reverseString(list[:shift])
	reverseString(list[shift:])
	reverseString(list)
	return string(list)
}

func reverseString(s []byte)  {
    start, end := 0, len(s)-1
    for start <= len(s)/2 -1 {
        s[start], s[end] = s[end], s[start]
        start++
        end--
    }
}

func main() {

	fmt.Println("1234567 " + " Rotate 3 : " + rotateString("1234567", 3))
	fmt.Println("1234567 " + " Rotate 4 : " + rotateString("1234567", 4))
	fmt.Println("1234567 " + " Rotate 5 : " + rotateString("1234567", 5))
	fmt.Println("1234567 " + " Rotate 6 : " + rotateString("1234567", 6))
	fmt.Println("1234567 " + " Rotate 7 : " + rotateString("1234567", 7))
	fmt.Println("1234567 " + " Rotate 8 : " + rotateString("1234567", 8))
}

