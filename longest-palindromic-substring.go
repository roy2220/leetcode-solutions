const N = 1000
var m [N][N]bool


func longestPalindrome(s string) string {
    max_i, max_j := 0, -1
    n := len(s)

    for i := n - 1; i >= 0; i-- {
        for j := n - 1; j >= 0; j-- {
            if (i >= j) {
                m[i][j] = true
            } else {
                m[i][j] = s[i] == s[j] && m[i + 1][j - 1]
            }

            if m[i][j] && j - i > max_j - max_i {
                max_i, max_j = i, j
            }
        }
    }

    return s[max_i:max_j + 1]
}
