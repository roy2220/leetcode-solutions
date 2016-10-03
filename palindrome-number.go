import (
    "strconv"
)


func isPalindrome(x int) bool {
    s := strconv.Itoa(x)
    return isPalindromicString(s)
}


func isPalindromicString(s string) bool {
    l := len(s)

    for i := 0; i <= (l - 1) / 2; i++ {
        j := (l - 1) - i

        if s[i] != s[j] {
            return false
        }
    }

    return true
}
