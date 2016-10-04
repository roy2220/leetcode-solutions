func nextPermutation(nums []int)  {
    var i int

    for i = len(nums) - 2; i >= 0; i-- {
        if nums[i] < nums[i + 1] {
            break
        }
    }

    if i < 0 {
        reverse(nums)
    } else {
        var j int
        for j = len(nums) - 1; nums[j] <= nums[i]; j-- {}
        swap(&nums[i], &nums[j])
        reverse(nums[i + 1:len(nums)])
    }
}


func swap(x *int, y *int) {
    z := *x
    *x = *y
    *y = z
}


func reverse(a []int) {
    l := len(a)

    for i := 0; i < l / 2; i++ {
        swap(&a[i], &a[l - 1 - i])
    }
}
