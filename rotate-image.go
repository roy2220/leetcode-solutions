func rotate(matrix [][]int)  {
    reserve1(matrix)
    reserve2(matrix)
}


func reserve1(matrix [][]int) {
    n := len(matrix)

    for y := 0; y < n - 1; y++ {
        for x := y + 1; x < n; x++ {
            swap(&matrix[y][x], &matrix[x][y])
        }
    }
}


func reserve2(matrix [][]int) {
    n := len(matrix)

    for y := 0; y < n; y++ {
        for x := 0; x < n / 2; x++ {
            swap(&matrix[y][x], &matrix[y][n - 1 - x])
        }
    }
}


func swap(x *int, y *int) {
    z := *x
    *x = *y
    *y = z
}
