func exist(board [][]byte, word string) bool {
    m := len(board)

    if m == 0 {
        return false
    }

    n := len(board[0])
    var dfs func (int, int, int) bool

    dfs = func (i int, x int, y int) bool {
        if i == len(word) {
            return true
        }

        if x < 0 || x >= m || y < 0 || y >= n || board[x][y] != word[i] {
            return false
        }

        as := [][]int{[]int{1,0},[]int{0,1},[]int{-1,0},[]int{0,-1}}
        board[x][y] = 0

        for _, a := range as {
            ok := dfs(i + 1, x + a[0], y + a[1])

            if ok {
                return true
            }
        }

        board[x][y] = word[i]
        return false
    }

    for x := 0; x < m; x++ {
        for y := 0; y < n; y++ {
            ok := dfs(0, x, y)

            if ok {
                return true
            }
        }
    }

    return false
}
