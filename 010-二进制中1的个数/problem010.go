package problem010


func Ones(n int) int {
    count := 0;
    for n!=0 {
        count++
        n = n & (n-1)
    }
    return count
}