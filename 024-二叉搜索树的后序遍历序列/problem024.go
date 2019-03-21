package main

import (
	"fmt"
)


func isPostOrder(post []int) bool {
	if len(post) <= 2 { return true }
	root := post[len(post)-1]
	left := -1
	for i:=len(post)-2; i>=0; i--{
		if post[i] < root {
			left = i
			break
		}
	}

	for _,v := range post[:left+1] {
		if v > root {
			return false
		}
	}

	return isPostOrder(post[0:left+1]) && isPostOrder(post[left+1:len(post)-1])

}

func main() {
	//test
	post1 := []int{ 2, 9, 5, 16, 17, 15, 19, 18, 12 }
	post2 := []int{ 2, 13, 5, 16, 17, 15, 19, 18, 12 }
	post3 := []int{ 2, 9, 5, 16, 17, 15, 11, 18, 12 }

	fmt.Println(post1 ,"is a postOrder ?  ", isPostOrder(post1))
	fmt.Println(post2 ,"is a postOrder ?  ", isPostOrder(post2))
	fmt.Println(post3 ,"is a postOrder ?  ", isPostOrder(post3))
	fmt.Printf("\n")
}