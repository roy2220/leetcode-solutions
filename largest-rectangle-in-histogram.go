func largestRectangleArea(heights []int) int {
    if len(heights) == 0 {
        return 0
    }

    q := [][]int{[]int{0, len(heights) - 1}}
    s := 0

    for len(q) >= 1 {
        i, j := q[0][0], q[0][1]
        q = q[1:]
        m := heights[i]

        for k := i + 1; k <= j; k++ {
            if heights[k] < m {
                m = heights[k]
            }
        }

        if s < (j - i + 1) * m {
            s = (j - i + 1) * m
        }

        for k := i; k <= j; k++ {
            if heights[k] == m {
                if k - 1 >= i {
                    q = append(q, []int{i, k - 1})
                }

                i = k + 1
            }
        }

        if i <= j {
            q = append(q, []int{i, j})
        }
    }

    return s
}
