func reverse(x int) int {
    var sign int

    if x < 0 {
        sign = -1
    } else {
        sign = 1
    }

    x /= sign
    y := 0

    for x >= 1 {
        if float64(y) > float64(int(^uint32(0) >> 1) - x % 10) / 10.0 {
            return 0
        }

        y = y * 10 + x % 10
        x /= 10
    }

    return sign * y
}
