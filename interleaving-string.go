func isInterleave(s1 string, s2 string, s3 string) bool {
    if len(s1) + len(s2) != len(s3) {
        return false
    }

    n := len(s1) + 1
    m := len(s2) + 1
    b := make([][]bool, n, n)

    for i := 0; i < n; i++ {
        b[i] = make([]bool, m, m)
    }

    b[0][0] = true

    for x := 1; x < n; x++ {
        b[x][0] = b[x - 1][0] && s1[x - 1] == s3[x - 1]
    }

    for y := 1; y < m; y++ {
        b[0][y] = b[0][y - 1] && s2[y - 1] == s3[y - 1]
    }

    for x := 1; x < n; x++ {
        for y := 1; y < m; y++ {
            b[x][y] = (s1[x - 1] == s3[x + y - 1] && b[x - 1][y]) ||
                      (s2[y - 1] == s3[x + y - 1] && b[x][y - 1])
        }
    }

    return b[len(s1)][len(s2)]
}

/*
func isInterleave(s1 string, s2 string, s3 string) bool {
    if !test(s1, s2, s3) {
        return false
    }

    if s1 + s2 == s3 || s2 + s1 == s3 {
        return true
    }

    l := len(s3) / 2
    i := len(s1)

    if i > l {
        i = l
    }

    j := l - i

    for i >= 0 && j <= len(s2) {
        f := isInterleave(s1[:i], s2[:j], s3[:l]) && isInterleave(s1[i:], s2[j:], s3[l:])

        if f {
            return true
        }

        i--
        j++
    }

    return false
}

func test(s1 string, s2 string, s3 string) bool {
    m := map[rune]int{}

    for _, c := range s1 {
        m[c]++
    }

    for _, c := range s2 {
        m[c]++
    }

    for _, c := range s3 {
        m[c]--
    }

    for _, v := range m {
        if v != 0 {
            return false
        }
    }

    return true
}
*/
