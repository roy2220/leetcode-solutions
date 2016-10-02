func lengthOfLongestSubstring(s string) int {
    return len(getLongestXSubstring(s))
}


// f(idx, set) = {
//     0                              (idx < 0 or str[idx] in set)
//     f(idx - 1, set + str[idx]) + 1 (idx >= 0 or str[idx] not in set)
// }
func getLongestXSubstring(s string) string {
    m := map[rune]int{}
    max_i := 0;
    max_j := -1;
    i := 0

    for j, c := range s {
        if k, ok := m[c]; ok {
            for l := i; l < k; l++ {
                delete(m, rune(s[l]))
            }

            i = k + 1
        }

        m[c] = j

        if j - i > max_j - max_i {
            max_i = i
            max_j = j
        }
    }

    return s[max_i:max_j + 1]
}
