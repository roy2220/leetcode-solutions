func searchRange(nums []int, target int) []int {
    start := bsearch_first(nums, 0, len(nums) - 1, target)
    end := bsearch_last(nums, 0, len(nums) - 1, target)
    return []int{start, end}
}


func bsearch_first(a []int, i int, j int, x int) int {
    for i < j {
        k := (i + j) / 2
        // i <= k < j

        if x <= a[k] {
            j = k
            // i <= j
        } else {
            i = k + 1
            // i <= j
        }
    }
    // i == j

    if a[i] == x {
        return i
    } else {
        return -1
    }
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
    } else {
        return -1
    }
}
