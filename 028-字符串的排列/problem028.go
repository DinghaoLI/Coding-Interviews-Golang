package main

import (
	"fmt"
)

func stringPermutation(str string) []string {
	var res []string
	s := []byte(str)
	length := len(s)
	var dfs func(idx int)
	dfs = func(idx int){
		if idx == length {
			str := string(s)
			res = append(res, str)
			return
		}
		m := make(map[byte]bool)
		for i:=idx; i<length; i++{
			if _,ok := m[s[i]]; ok { continue }
			m[s[i]] = true
			s[idx], s[i] = s[i], s[idx]
			dfs(idx+1)
			s[i], s[idx] = s[idx], s[i]
		}
	}
	dfs(0)
	return res
}

func main() {
	//test
	s := "abc"
	b := "ABC"
	d := "aab"
	fmt.Println(s, " Permutation: ", stringPermutation(s))
	fmt.Println(b, " Permutation: ", stringPermutation(b))
    fmt.Println(d, " Permutation: ", stringPermutation(d))

}
