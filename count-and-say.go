import (
    "fmt"
)


func countAndSay(n int) string {
    r := "1"

    for n >= 2 {
        r = f(r)
        n--
    }

    return r
}


func f(s string) string {
    r := ""
    j := 0
    pc := s[0]

    for i := 1; i < len(s); i++ {
        if s[i] != pc {
            n := i - j
            r += fmt.Sprintf("%d%c", n, pc)
            j = i
            pc = s[i]
        }

    }

    r += fmt.Sprintf("%d%c", len(s) - j, pc)
    return r
}
