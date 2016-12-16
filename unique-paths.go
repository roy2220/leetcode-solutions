func uniquePaths(m int, n int) int {
    g := make([][]int, m, m)

    for i := 0; i < m; i++ {
        g[i] = make([]int, n, n)
    }

    g[m - 1][n - 1] = 1

    for x := 0; x < m - 1; x++ {
        g[x][n - 1] = 1
    }

    for y := 0; y < n - 1; y++ {
        g[m - 1][y] = 1
    }

    for x := m - 2; x >= 0; x-- {
        for y := n - 2; y >= 0; y-- {
            g[x][y] = g[x + 1][y] + g[x][y + 1]
        }
    }

    return g[0][0]
}
