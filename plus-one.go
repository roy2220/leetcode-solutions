func plusOne(digits []int) []int {
    r := []int{}
    a := 1

    for i := len(digits) - 1; i >= 0; i-- {
        k := digits[i] + a
        r = append([]int{k % 10}, r...)
        a = k / 10
    }

    if a >= 1 {
        r = append([]int{a}, r...)
    }

    return r
}
