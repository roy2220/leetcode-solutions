func sortColors(nums []int)  {
    i := 0
    j := len(nums) - 1
    k := i

    for k <= j {
        x := nums[k]

        if x == 0 {
            if k == i {
                k++
            } else {
                nums[k], nums[i] = nums[i], nums[k]
            }

            i++
        } else if x == 2 {
            nums[k], nums[j] = nums[j], nums[k]
            j--
        } else {
            k++
        }
    }
}
