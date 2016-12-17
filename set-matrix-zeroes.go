func setZeroes(matrix [][]int)  {
    if len(matrix) == 0 {
        return
    }

    for x := 0; x < len(matrix) - 1; x++ {
        f := false

        for y := 0; y < len(matrix[0]); y++ {
            if matrix[x][y] == 0 && (x == 0 || matrix[x - 1][y] != 0) {
                f = true
                break
            }
        }

        for y := 0; y < len(matrix[0]); y++ {
            if matrix[x + 1][y] == 0 {
                if matrix[x][y] == 0 {
                    matrix[x][y] = 1
                }
            } else {
                if matrix[x][y] == 0 {
                    matrix[x + 1][y] = 0
                }

                if f {
                    matrix[x][y] = 0
                }
            }
        }
    }

    x := len(matrix) - 1
    f := false

    for y := 0; y < len(matrix[0]); y++ {
        if matrix[x][y] == 0 {
            if x == 0 || matrix[x - 1][y] != 0 {
                f = true
            }

            for z := 0; z < x; z++ {
                matrix[z][y] = 0
            }
        }
    }

    if f {
        for y := 0; y < len(matrix[0]); y++ {
            matrix[x][y] = 0
        }
    }
}
