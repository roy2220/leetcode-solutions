func searchMatrix(matrix [][]int, target int) bool {
    m := len(matrix)

    if m == 0 {
        return false
    }

    n := len(matrix[0])

    get_matrix := func (i int) int {
        return matrix[i / n][i % n]
    }

    lo := 0
    hi := m * n - 1

    for lo <= hi {
        mi := (lo + hi) / 2
        x := get_matrix(mi)

        if target < x {
            hi = mi - 1
        } else if target > x {
            lo = mi + 1
        } else {
            return true
        }
    }

    return false
}
