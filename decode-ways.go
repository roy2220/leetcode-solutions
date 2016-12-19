func numDecodings(s string) int {
    if len(s) == 0 {
        return 0
    }

    var dfs func (int) int

    dfs = func (i int) int {
        if i == len(s) {
            return 1
        }

        if s[i] == '0' {
            return 0
        }

        c := dfs(i + 1)

        if i + 1 < len(s) && (s[i] == '1' || (s[i] == '2' && s[i + 1] <= '6')) {
            c += dfs(i + 2)
        }

        return c
    }

    return dfs(0)
}
