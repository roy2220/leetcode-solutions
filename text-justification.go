import "strings"

func fullJustify(words []string, maxWidth int) []string {
    r := []string{}
    q := []string{}
    ql := -1

    for _, w := range words {
        if ql + 1 + len(w) > maxWidth {
            r = append(r, makeLine(q, ql, maxWidth))
            q = []string{w}
            ql = len(w)
        } else {
            q = append(q, w)
            ql += 1 + len(w)
        }
    }

    if len(q) >= 1 {
        r = append(r, makeLastLine(q, ql, maxWidth))
    }

    return r
}

func makeLine(q []string, ql int, maxWidth int) string {
    if len(q) == 1 {
        return q[0] + strings.Repeat(" ", maxWidth - ql)
    } else {
        l1 := maxWidth - ql
        m := len(q) - 1
        n := l1 / m
        l2 := m * n
        r := q[0]

        for i := 1; i < len(q); i++ {
            if l1 > l2 {
                r += strings.Repeat(" ", n + 2) + q[i]
                l1 -= n + 1
            } else {
                r += strings.Repeat(" ", n + 1) + q[i]
                l1 -= n
            }

            l2 -= n
        }

        return r
    }
}

func makeLastLine(q []string, ql int, maxWidth int) string {
    return strings.Join(q, " ") + strings.Repeat(" ", maxWidth - ql)
}
