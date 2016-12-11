import "sort"

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
    return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
    return len(s)
}

func SortString(s string) string {
    r := []rune(s)
    sort.Sort(sortRunes(r))
    return string(r)
}

func groupAnagrams(strs []string) [][]string {
    m := map[string][]string{}

    for _, str := range strs {
        sorted_str := SortString(str)
        var arr []string
        arr = m[sorted_str]

        if arr == nil {
            arr = []string{}
        }

        m[sorted_str] = append(arr, str)
    }

    r := [][]string{}

    for _, v := m {
        r = append(r, v)
    }

    return r
}
