package main

import (
	"fmt"
)

func cycle(n,m int) int {
	if n < 1 || m < 1 {
        return -1
    } else if n == 1{
        return 0
    } else {
        // F[n] = (F[n - 1] + m) % n
        return (cycle(n-1,m) + m) % n
    }
}

func main() {
	
    fmt.Println(3,2, " => ", cycle(3,2))

    fmt.Println(4,2, " => ", cycle(4,2))

    fmt.Println(5,2, " => ", cycle(5,2))
}
