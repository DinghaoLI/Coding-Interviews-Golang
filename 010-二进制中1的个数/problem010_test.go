package problem010

import (
    "testing"
)

func Test_case(t *testing.T) {
    if Ones(0) == 0 {
        t.Log("Pass")
    } else {
        t.Error("Failed")
    }
    // int64
    if Ones(-1) == 64 {
        t.Log("Pass")
    } else {
        t.Error("Failed value: ",Ones(-1))
    }

    if Ones(8) == 1 {
        t.Log("Pass")
    } else {
        t.Error("Failed")
    }


    if Ones(5) == 2 {
        t.Log("Pass")
    } else {
        t.Error("Failed")
    }

}

