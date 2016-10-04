func letterCombinations(digits string) []string {
    if len(digits) == 0 {
        return []string{}
    }

    tab := []string{
        " ",
        "",
        "abc",
        "def",
        "ghi",
        "jkl",
        "mno",
        "pqrs",
        "tuv",
        "wxyz",
    }

    res := []string{""}

    for _, dig := range digits {
        keys := tab[int(dig - '0')]
        new_res := []string{}

        for _, key := range keys {
            for _, str := range res {
                str += string(key)
                new_res = append(new_res, str)
            }
        }

        res = new_res
    }

    return res
}
