package main

import (
	"fmt"
)

func count(str string) int {
	bytes := []byte(str)
	m := make(map[byte]int)
	for _,v := range bytes {
		m[v]++
	}
	for i,v := range bytes {
		if m[v] == 1 {
			return i
		}
	}
	return -1
}


func main() {
	fmt.Println("alskdphkjsbcjkvhaklsuhdtq: ",count("alskdphkjsbcjkvhaklsuhdtq"))
	fmt.Println("aaa2bbbfff3gggdddsssr3rrwwsss4: ",count("aaa2bbbfff3gggdddsssrrrwwsss4"))

}
















