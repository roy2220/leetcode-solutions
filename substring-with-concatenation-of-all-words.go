import (
    "strings"
)


func findSubstring(s string, words []string) []int {
    r := []int{}
    d1 := map[string]int{}

    for _, w := range words {
        dict_add_word(d1, w)
    }

    wl := len(words[0])
    chains := split_into_chains(s, wl)

    for i, chain := range chains {
        if len(chain) >= len(words) {
            d2 := map[string]int{}

            for j := 0; j < len(words); j++ {
                dict_add_word(d2, chain[j])
            }

            if dict_equal(d1, d2) {
                r = append(r, i)
            }

            for j := len(words); j < len(chain); j++ {
                k := j - len(words)
                dict_del_word(d2, chain[k])
                dict_add_word(d2, chain[j])

                if dict_equal(d1, d2) {
                    r = append(r, i + (k + 1) * wl)
                }
            }
        }
    }

    return r
}


func dict_add_word(d map[string]int, w string) {
    d[w] += 1
}


func dict_del_word(d map[string]int, w string) {
    d[w] -= 1
}


func dict_equal(d1 map[string]int, d2 map[string]int) bool {
    for w, n := range d1 {
        if d2[w] != n {
            return false
        }
    }

    return true
}


func split_into_chains(s string, wl int) [][]string {
    r := [][]string{}

    if len(s) < wl {
        return r
    }

    for i := 0; i < wl; i++ {
        c := []string{}

        for j := i + wl; j <= len(s); j += wl {
            c = append(c, s[j - wl:j])
        }

        r = append(r, c)
    }

    return r
}
