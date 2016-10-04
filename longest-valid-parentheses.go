func longestValidParentheses(s string) int {
    r := 0
    var f func (int) int

    f = func (i int) int {
        j := i

        for j < len(s) && s[j] != ')' {
            j = f(j + 1) + 1

            if j <= len(s) {
                l := j - i

                if l > r {
                    r = l
                }
            }
        }

        return j
    }

    i := f(0)

    for i < len(s) {
        i = f(i + 1)
    }

    return r
}
