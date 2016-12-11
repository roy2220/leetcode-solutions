func jump(nums []int) int {
    i := -1
    j := 0
    s := 0

    for j < len(nums) - 1 {
        new_j := j

        for k := i + 1; k <= j; k++ {
            if k + nums[k] > new_j {
                i = k
                new_j = k + nums[k]
            }
        }

        if new_j == j {
            return -1
        }

        j = new_j
        s++
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
