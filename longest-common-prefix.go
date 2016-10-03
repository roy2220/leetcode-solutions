func longestCommonPrefix(strs []string) string {
    if len(strs) == 0 {
        return ""
    }

    for i := 0;; i++ {
        str := strs[0]

        if len(str) <= i {
            return str[0:i]
        }

        c := str[i]

        for _, str := range strs[1:len(strs)] {
            if len(str) <= i || str[i] != c {
                return str[0:i]
            }
        }
    }
}
