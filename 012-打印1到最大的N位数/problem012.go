package main

import (
    "fmt"
    "errors"
)

var MaxInt int = int(^uint(0) >> 1)

// problem 011
func Pow(base float64, exp int) (float64, error) {
    if exp == 0 {
        return 1, nil
    }
    if exp < 0 && base == 0 {
        return -1 , errors.New("base == 0 and exp < 0")
    }

    if exp > 0 {
        return PowNormal(base, exp), nil    
    } else {
        res := PowNormal(base, -exp)
        res = 1/res
        return res, nil
    }   

}

func PowNormal(base float64, exp int) float64 {
    res, temp := 1.0, base
    for exp != 0 {
        if exp&1 == 1 {
            res *= temp
        }
        temp *= temp
        exp >>= 1
    }
    return res
}

// 012
func showN(n int) error {
    top, _ := Pow(10, n)

    if (int(top) < 0){
        return errors.New("Overflow")
    }

    for i := 1; i < int(top); i++ {
        fmt.Println(i)
    }
    return nil 
}

func main() {
    showN(1)
    showN(3)
}




















