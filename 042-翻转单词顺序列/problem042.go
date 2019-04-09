package main

import (
	"fmt"
)

func reverseWord(str string) string {
	list := []byte(str)
	reverseString(list)
	left, right := 0, 0
	for right<=len(list)-1 {
		if list[right] != ' ' {
			right++
			continue
		}
		reverseString(list[left:right]) 
		left, right = right+1, right+1
	}
	reverseString(list[left:right]) 
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

	fmt.Println("student. a am I" + "  ===》  " + reverseWord("student. a am I"))
	fmt.Println("I'm a Freshman and I like JOBDU!" +  "  ===》  " + reverseWord("I'm a Freshman and I like JOBDU!"))

}

