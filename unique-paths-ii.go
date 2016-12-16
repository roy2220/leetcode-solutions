func uniquePathsWithObstacles(obstacleGrid [][]int) int {
    m := len(obstacleGrid)
    n := len(obstacleGrid[0])
    g := make([][]int, m, m)

    for i := 0; i < m; i++ {
        g[i] = make([]int, n, n)
    }

    if obstacleGrid[m - 1][n - 1] == 0 {
        g[m - 1][n - 1] = 1
    } else {
        g[m - 1][n - 1] = 0
    }

    for x := m - 2; x >= 0; x-- {
        if obstacleGrid[x][n - 1] == 0 {
            g[x][n - 1] = g[x + 1][n - 1]
        } else {
            g[x][n - 1] = 0
        }
    }

    for y := n - 2; y >= 0; y-- {
        if obstacleGrid[m - 1][y] == 0 {
            g[m - 1][y] = g[m - 1][y + 1]
        } else {
            g[m - 1][y] = 0
        }
    }

    for x := m - 2; x >= 0; x-- {
        for y := n - 2; y >= 0; y-- {
            if obstacleGrid[x][y] == 0 {
                g[x][y] = g[x + 1][y] + g[x][y + 1]
            } else {
                g[x][y] = 0
            }
        }
    }

    return g[0][0]
}
