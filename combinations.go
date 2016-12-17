func combine(n int, k int) [][]int {
    c := []int{}
    r := [][]int{}
    var dfs func (i int)

    dfs = func (i int) {
        if len(c) == k {
            c2 := make([]int, k)
            copy(c2, c)
            r = append(r, c2)
            return
        }

        for ; i <= n - (k - len(c)) + 1; i++ {
            c = append(c, i)
            dfs(i + 1)
            c = c[:len(c) - 1]
        }
    }

    dfs(1)
    return r
}
