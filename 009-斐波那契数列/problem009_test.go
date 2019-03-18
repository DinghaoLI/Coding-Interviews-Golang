package problem009

import (
    "testing"
)

func Test_case(t *testing.T) {
    if Fibonacci(0) == 0 {
        t.Log("Pass")
    } else {
        t.Error("Failed")
    }

    if Fibonacci(1) == 1 {
        t.Log("Pass")
    } else {
        t.Error("Failed")
    }

    if Fibonacci(7) == 13 {
        t.Log("Pass")
    } else {
        t.Error("Failed")
    }


    if Fibonacci(13) == 233 {
        t.Log("Pass")
    } else {
        t.Error("Failed")
    }

}

