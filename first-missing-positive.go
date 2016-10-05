func firstMissingPositive(nums []int) int {
    i := 0
    j := len(nums)

    for i < j {
        k := nums[i] - 1

        if k == i {
            i++
        } else if k > i && k < j && nums[i] != nums[k] {
            swap(&nums[i], &nums[k])
        } else {
            swap(&nums[i], &nums[j - 1])
            j--
        }
    }

    return i + 1
}


func swap(x *int, y *int) {
    z := *x
    *x = *y
    *y = z
}
