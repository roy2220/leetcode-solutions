func addBinary(a string, b string) string {
    r := ""
    x := 0

    for i, j := len(a) - 1, len(b) - 1; i >= 0 || j >= 0; i, j = i - 1, j - 1 {
        var v1, v2 int

        if i >= 0 {
            v1 = int(a[i] - '0')
        } else {
            v1 = 0
        }

        if j >= 0 {
            v2 = int(b[j] - '0')
        } else {
            v2 = 0
        }

        k := v1 + v2 + x
        r = string('0' + k % 2) + r
        x = k / 2
    }

    if x >= 1 {
        r = string('0' + x) + r
    }

    return r
}
