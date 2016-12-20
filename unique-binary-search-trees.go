func numTrees(n int) int {
    c := []int{1,1}

    for i := 2; i <= n; i++ {
        n := 0

        for j := 0; j < i; j++ {
            n += c[j] * c[i - 1 - j]
        }

        c = append(c, n)
    }

    return c[n]
}

/*
func numTrees(n int) int {
    m := map[string]int{}
    var f func (int, int) int

    f = func (i int, j int) int {
        if i >= j {
            return 1
        }

        key := fmt.Sprintf("%d,%d", i, j)

        if m[key] != 0 {
            return m[key]
        }

        n := 0

        for k := i; k <= j; k++ {
            n += f(i, k - 1) * f(k + 1, j)
        }

        m[key] = n
        return n
    }

    return f(1, n)
}
*/
