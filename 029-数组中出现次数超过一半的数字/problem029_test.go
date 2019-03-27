package problem010

import (
    "testing"
)

func Test_case(t *testing.T) {
   nums := []int{2,1,2,1,2}
   res, err := getMostFreq(nums)
   if res == 2 {
        t.Log("Pass")
    } else {
        t.Error("Failed")
    }

   nums = []int{1,2,2,2,2,6,7}
   res, err = getMostFreq(nums)
   if res == 2 {
        t.Log("Pass")
    } else {
        t.Error("Failed")
    }

   nums = []int{1}
   res, err = getMostFreq(nums)
   if res == 1 {
        t.Log("Pass")
    } else {
        t.Error("Failed")
    }

   nums = []int{}
   res, err = getMostFreq(nums)
   if err != nil {
        t.Log("Pass")
    } else {
        t.Error("Failed")
    }

}

