func isScramble(s1 string, s2 string) bool {
    if len(s1) == 1 {
        return s1 == s2
    } else if len(s1) == 2 {
        return s1 == s2 || (s1[0] == s2[1] && s1[1] == s2[0])
    } else {
        if !test(s1, s2) {
            return false
        }

        for l1 := 1; l1 < len(s1); l1++ {
            l2 := len(s1) - l1

            f := (isScramble(s1[:l1], s2[:l1]) && isScramble(s1[l1:], s2[l1:])) ||
                 (isScramble(s1[:l1], s2[l2:]) && isScramble(s1[l1:], s2[:l2]))

            if f {
                return true
            }
        }

        return false
    }
}

func test(s1 string, s2 string) bool {
    m := map[rune]int{}

    for _, c := range s1 {
        m[c]++
    }

    for _, c := range s2 {
        m[c]--
    }

    for _, v := range m {
        if v != 0 {
            return false
        }
    }

    return true
}
