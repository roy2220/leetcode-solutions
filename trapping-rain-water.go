func trap(height []int) int {
    r := 0
    i := 0
    j := len(height) - 1

    for i < j {
        if height[i] < height[j] {
            h := height[i]

            for i++; i < j; i++ {
                if height[i] > h {
                    break
                }

                r += h - height[i]
            }
        } else {
            h := height[j]

            for j--; j > i; j-- {
                if height[j] > h {
                    break
                }

                r += h - height[j]
            }
        }
    }

    return r
}
