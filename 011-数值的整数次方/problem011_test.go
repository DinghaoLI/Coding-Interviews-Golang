package problem011

import (
    "testing"
)

func Test_case(t *testing.T) {
    res, _ := Pow(0, 0)
    if res == 1 {
        t.Log("Pass")
    } else {
        t.Error("Failed")
    }
    
    res, _ = Pow(2, -1)
    if res == 0.5 {
        t.Log("Pass")
    } else {
        t.Error("Failed value: ")
    }

    res, _ = Pow(2, 3)
    if res == 8 {
        t.Log("Pass")
    } else {
        t.Error("Failed")
    }

    res, _ = Pow(2, -2)
    if res == 0.25 {
        t.Log("Pass")
    } else {
        t.Error("Failed")
    }

    res, err := Pow(0, -2)
    if err != nil  {
        t.Log("Pass")
    } else {
        t.Error("Failed")
    }

}

