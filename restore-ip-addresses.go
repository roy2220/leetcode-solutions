import "strings"

func restoreIpAddresses(s string) []string {
    r := []string{}
    ip := []string{}
    var dfs func (int)

    dfs = func (i int) {
        if len(ip) == 4 {
            if i == len(s) {
                r = append(r, strings.Join(ip, "."))
            }

            return
        }

        n := 0

        for j := i; j < len(s); j++ {
            n = 10 * n + int(s[j] - '0')

            if j - i + 1 >= 2 && s[i] == '0' {
                break
            }

            if n >= 256 {
                break
            }

            ip = append(ip, s[i:j + 1])
            dfs(j + 1)
            ip = ip[:len(ip) - 1]
        }
    }

    dfs(0)
    return r
}
