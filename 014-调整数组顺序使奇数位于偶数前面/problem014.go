package main

import (
    "fmt"
)


func oddFirst(s []int){
    left, right := 0, len(s)-1
    for left < right {
        for s[right]%2==0 && left<right{
            right--
        }
        for s[left]%2==1 && left<right{
            left++
        }
        if left==right {
            break
        }
        if left<right {
            s[left], s[right] = s[right], s[left]
        }
    }
}



func main() {
    //test

    l := []int{1,2,3,4,5,6,7}
    fmt.Println("before: ", l)
    oddFirst(l)
    fmt.Println("after: ",l)

    l = []int{2,2,2,2,5,6,7}
    fmt.Println("before: ", l)
    oddFirst(l)
    fmt.Println("after: ",l)
}





