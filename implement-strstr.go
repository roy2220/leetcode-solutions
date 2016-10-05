func strStr(haystack string, needle string) int {
    var i int

    for i = 0; i <= len(haystack) - len(needle); i++ {
        var j int

        for j = 0; j < len(needle); j++ {
            if haystack[i + j] != needle[j] {
                break
            }
        }

        if j == len(needle) {
            break
        }
    }

    if i <= len(haystack) - len(needle) {
        return i
    } else {
        return -1
    }
}
