func searchInsert(nums []int, target int) int {
    return bsearch_last(nums, 0, len(nums) - 1, target)
}


func bsearch_last(a []int, i int, j int, x int) int {
    for i < j {
        k := (i + j + 1) / 2
        // i < k <= j

        if x >= a[k] {
            i = k
            // i <= j
        } else {
            j = k - 1
            // i <= j
        }
    }
    // i == j

    if a[i] == x {
        return i
    } else if x < a[i] {
        return i
    } else {
        return i + 1
    }
}
