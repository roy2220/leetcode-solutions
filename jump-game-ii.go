func jump(nums []int) int {
    i := 0
    j := -1
    k := 0
    s := 0

    for i = k + 1; i < len(nums); i = k + 1 {
        s++

        for l := j + 1; l < i; l++ {
            if l + nums[l] > k {
                j = l
                k = l + nums[l]
            }
        }
    }

    return s
}


// func jump(nums []int) int {
//     a := make([]int, len(nums))
//     a[len(nums) - 1] = 0

//     for i := len(nums) - 2; i >= 0; i-- {
//         n := min(i + nums[i] + 1, len(nums))

//         j := i + 1
//         k := j

//         for j++; j < n; j++ {
//             if a[j] < a[k] {
//                 k = j
//             }
//         }

//         a[i] = a[k] + 1
//     }

//     return a[0]
// }


// func min(x int, y int) int {
//     if x < y {
//         return x
//     } else {
//         return y
//     }
// }
