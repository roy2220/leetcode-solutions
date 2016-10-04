func longestValidParentheses(s string) int {
    r := 0
    var f func (int) int

    f = func (i int) int {
        j := i

        for {
            if j >= len(s) || s[j] == ')' {
                return j
            } else {
                j = f(j + 1) + 1

                if j <= len(s) {
                    l := j - i

                    if l > r {
                        r = l
                    }
                }
            }
        }
    }

    i := f(0)

    for i < len(s) {
        i = f(i + 1)
    }

    return r
}
