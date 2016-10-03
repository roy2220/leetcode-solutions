func maxArea(height []int) int {
    return f(height, 0, len(height) - 1)
}


func f(a []int, i int, j int) int {
    if i >= j {
        return 0
    } else {
        x, y := a[i], a[j]

        if x < y {
            s := (j - i) *  x
            return max(s, f(a, i + 1, j))
        } else {
            s := (j - i) *  y
            return max(s, f(a, i, j - 1))
        }
    }
}


func max(x int, y int) int {
    if x > y {
        return x
    } else {
        return y
    }
}
