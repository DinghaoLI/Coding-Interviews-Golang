package problem008

import (
    "testing"
)

func Test_case1(t *testing.T) {
    if minNumberInRotateArray([]int{3, 4, 5, 1, 2}) == 1 {
        t.Log("Pass")
    } else {
        t.Error("Failed")
    }
}

func Test_case2(t *testing.T) {
  if minNumberInRotateArray([]int{1 , 0, 1, 1, 1}) == 0 {
        t.Log("Pass")
    } else {
        t.Error("Failed")
    }
}

func Test_case3(t *testing.T) {
  if minNumberInRotateArray([]int{2,2,3,4,1,1,2,2,2}) == 1 {
        t.Log("Pass")
    } else {
        t.Error("Failed")
    }
}