func climbStairs(n int) int {
    if n < 2 {
        return n
    }

    x := 1
    y := 1

    for i := 1; i < n; i++ {
        z := x + y
        x = y
        y = z
    }

    return y
}
