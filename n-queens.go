func solveNQueens(n int) [][]string {
    m := make([][]bool, n, n)

    for i, _ := range m {
        m[i] = make([]bool, n, n)
    }

    r := [][]string{}
    dfs(m, 0, &r)
    return r
}

func isOnBoard(n int, x int, y int) bool {
    return x >= 0 && x < n && y >= 0 && y < n
}

func test(m [][]bool, x int, y int) bool {
    as := [][2]int {
        [2]int{-1, -1},
        [2]int{-1,  0},
        [2]int{-1,  1},
    }

    for _, a := range as {
        x1 := x + a[0]
        y1 := y + a[1]

        for isOnBoard(len(m), x1, y1) {
            if m[x1][y1] {
                return false
            }

            x1 += a[0]
            y1 += a[1]
        }
    }

    return true
}

func dfs(m [][]bool, i int, r *[][]string) {
    if i == len(m) {
        result := []string{}

        for _, v1 := range m {
            s := ""

            for _, v2 := range v1 {
                if v2 {
                    s += "Q"
                } else {
                    s += "."
                }
            }

            result = append(result, s)
        }

        *r = append(*r, result)
        return
    }

    for j := 0; j < len(m); j++ {
        if test(m, i, j) {
            m[i][j] = true
            dfs(m, i + 1, r)
            m[i][j] = false
        }
    }
}
