func removeDuplicates(nums []int) int {
    if len(nums) == 0 {
        return 0
    }

    i := 0
    var n int
    pnum := ^nums[0]

    for _, num := range nums {
        if num == pnum {
            n++
        } else {
            n = 1
        }

        if n <= 2 {
            nums[i] = num
            i++
        }

        pnum = num
    }

    return i
}

func min(x int, y int) int {
    if x < y {
        return x
    } else {
        return y
    }
}
