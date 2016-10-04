func removeElement(nums []int, val int) int {
    l := 0

    for _, x := range nums {
        if x != val {
            nums[l] = x
            l++
        }
    }

    return l
}
