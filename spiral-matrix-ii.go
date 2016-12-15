func generateMatrix(n int) [][]int {
    result := make([][]int, n, n)

    if n == 0 {
        return result
    }

    for m := 0; m < n; m++ {
        result[m] = make([]int, n, n)
    }

    steps := [4][2]int{[2]int{0, 1}, [2]int{1, 0}, [2]int{0, -1}, [2]int{-1, 0}}
    stops := [4][2]int{[2]int{0, n - 1}, [2]int{n - 1, n - 1}, [2]int{n - 1, 0}, [2]int{1, 0}}
    x := 0
    y := -1
    z := 0
    i := 0

    for {
        if x == stops[i][0] && y == stops[i][1] {
            return result
        }

        for {
            x += steps[i][0]
            y += steps[i][1]
            z++
            result[x][y] = z

            if x == stops[i][0] && y == stops[i][1] {
                break
            }
        }

        j := (i + 1) % 4
        stops[i][0] += steps[j][0] - steps[i][0]
        stops[i][1] += steps[j][1] - steps[i][1]
        i = j
    }
}
