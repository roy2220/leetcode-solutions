func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
    if len(nums1) == 0 || len(nums2) == 0 {
        return [][]int{}
    }

    pairs := [][2]int{[2]int{0, 0}}
    result := [][]int{}

    for ;; {
        x, y := popSmallestPair(&pairs, nums1, nums2)
        result = append(result, []int{nums1[x], nums2[y]})

        if len(result) == k || (x + 1 == len(nums1) && y + 1 == len(nums2)) {
            break
        }

        if x + 1 < len(nums1) {
            addPair(&pairs, x + 1, y)
        }

        if y + 1 < len(nums2) {
            addPair(&pairs, x, y + 1)
        }
    }

    return result
}

func addPair(pairs *[][2]int, x int, y int) {
    for _, pair := range *pairs {
        if x >= pair[0] && y >= pair[1] {
            return
        }
    }

    *pairs = append(*pairs, [2]int{x, y})
}


func popSmallestPair(pairs *[][2]int, nums1 []int, nums2 []int) (int, int) {
    smallestPair := (*pairs)[0]

    for i := 1; i < len(*pairs); i++ {
        pair := (*pairs)[i]

        if nums1[pair[0]] + nums2[pair[1]] < nums1[smallestPair[0]] + nums2[smallestPair[1]] {
            (*pairs)[i - 1] = smallestPair
            smallestPair = pair
        } else {
            (*pairs)[i - 1] = pair
        }
    }

    *pairs = (*pairs)[:len(*pairs) - 1]
    return smallestPair[0], smallestPair[1]
}
