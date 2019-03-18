package problem009

func Fibonacci(n int) int {
    if n <= 1 {
        return n
    }
    f1, f2 := 0, 1
    res := 0
    for i := 2; i <= n; i++ {
    	res = f1+f2
    	f1, f2 = f2, res
    }
    return res
}
