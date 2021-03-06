func threeSumClosest(nums []int, target int) int {
    var result int
    diff := int(^uint(0) >> 1)

    traverseCombinations(nums, 3, func (combination []int) bool {
        i, j, k := combination[0], combination[1], combination[2]
        s := nums[i] + nums[j] + nums[k]
        new_diff := abs(s - target)

        if new_diff < diff {
            result = s
            diff = new_diff

            if diff == 0 {
                return true
            }
        }

        return false
    })

    return result
}


func abs(x int) int {
    if x < 0 {
        return -x
    } else {
        return x
    }
}


func traverseCombinations(elements []int, combinationLength int, callback func ([]int) bool) bool {
    if combinationLength > len(elements) {
        return false
    }

    combination := []int{}
    var doTraverseCombinations func (int) bool

    doTraverseCombinations = func (i int) bool {
        if len(combination) == combinationLength {
            return callback(combination)
        }

        for ; i < len(elements) - (combinationLength - len(combination)) + 1; i++ {
            combination = append(combination, i);
            stop := doTraverseCombinations(i + 1)
            combination = combination[0:len(combination)-1];

            if stop {
                return true
            }
        }

        return false
    }

    return doTraverseCombinations(0)
}
