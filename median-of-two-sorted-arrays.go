// NOT a correct answer, O(n) but passed through
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    nums3 := append(nums1, nums2...)
    l := len(nums3)

    if l % 2 == 1 {
        return float64(quickFindKth(nums3, 0, l - 1, l / 2))
    } else {
        f1 := float64(quickFindKth(nums3, 0, l - 1, l / 2 - 1))
        f2 := float64(quickFindKth(nums3, 0, l - 1, l / 2))
        return (f1 + f2) / 2
    }
}


func quickFindKth(a []int, i int, j int, k int) int {
    k2 := partition(a, i, j)

    if k2 == k {
        return a[k]
    } else if k2 < k {
        return quickFindKth(a, k2 + 1, j, k)
    } else {
        return quickFindKth(a, i, k2 - 1, k)
    }
}


func partition(a []int, i int, j int) int {
    k := i

    for ; i < j; i++ {
        if a[i] < a[j] {
            swap(&a[i], &a[k])
            k++
        }
    }

    swap(&a[j], &a[k])
    return k
}


func swap(x *int, y *int) {
    z := *x
    *x = *y
    *y = z
}
