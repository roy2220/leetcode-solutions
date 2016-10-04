func removeDuplicates(nums []int) int {
    if len(nums) == 0 {
        return 0
    }

    l := 0
    x := ^nums[0]

    for _, y := range nums {
        if y != x {
            nums[l] = y
            x = y
            l++
        }
    }

    return l
}
