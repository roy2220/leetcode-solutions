func minDistance(word1 string, word2 string) int {
    m := len(word1) + 1
    n := len(word2) + 1
    s := make([][]int, m, m)

    for i := 0; i < m; i++ {
        s[i] = make([]int, n, n)
    }

    s[0][0] = 0

    for x := 1; x < m; x++ {
        s[x][0] = x
    }

    for y := 1; y < n; y++ {
        s[0][y] = y
    }

    for x := 1; x < m; x++ {
        for y := 1; y < n; y++ {
            if word1[x - 1] == word2[y - 1] {
                s[x][y] = s[x - 1][y - 1]
            } else {
                min := s[x - 1][y - 1]

                if s[x][y - 1] < min {
                    min = s[x][y - 1]
                }

                if s[x - 1][y] < min {
                    min = s[x - 1][y]
                }

                s[x][y] = min + 1
            }
        }
    }

    return s[m - 1][n - 1]
}
