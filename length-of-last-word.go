func lengthOfLastWord(s string) int {
    l := 0
    el := 0

    for _, c := range s {
        if c == ' ' {
            if l >= 1 {
                el = l
            }

            l = 0
        } else {
            l++
        }
    }

    if l == 0 {
        l = el
    }

    return l
}
