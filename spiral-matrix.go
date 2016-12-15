func spiralOrder(matrix [][]int) []int {
    result := []int{}
    m := len(matrix)

    if m == 0 {
        return result
    }

    n := len(matrix[0])
    steps := [4][2]int{[2]int{0, 1}, [2]int{1, 0}, [2]int{0, -1}, [2]int{-1, 0}}
    stops := [4][2]int{[2]int{0, n - 1}, [2]int{m - 1, n - 1}, [2]int{m - 1, 0}, [2]int{1, 0}}
    x := 0
    y := -1
    i := 0

    for {
        if x == stops[i][0] && y == stops[i][1] {
            return result
        }

        for {
            x += steps[i][0]
            y += steps[i][1]
            result = append(result, matrix[x][y])

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
