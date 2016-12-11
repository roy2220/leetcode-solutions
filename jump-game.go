func canJump(nums []int) bool {
    i := -1
    j := 0

    for j < len(nums) - 1 {
        new_j := j

        for k := i + 1; k <= j; k++ {
            if k + nums[k] > new_j {
                i = k
                new_j = k + nums[k]
            }
        }

        if new_j == j {
            return false
        }

        j = new_j
    }

    return true
}
