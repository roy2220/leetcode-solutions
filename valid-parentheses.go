func isValid(s string) bool {
    pair := map[rune]rune{'}': '{', ')': '(', ']': '['}
    stk := []rune{'#'}

    for _, c := range s {
        if c2, ok := pair[c]; ok {
            if stk[len(stk) - 1] != c2 {
                return false
            }

            stk = stk[0:len(stk) - 1]
        } else {
            stk = append(stk, c)
        }
    }

    return len(stk) == 1
}
