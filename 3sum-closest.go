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
    if len(elements) < combinationLength {
        return false
    }

    elementIndex := 0
    combination := []int{}
    var doTraverseCombinations func () bool

    doTraverseCombinations = func () bool {
        if len(combination) < combinationLength {
            combination = append(combination, elementIndex)
            var stop bool

            if len(combination) == combinationLength {
                stop = callback(combination)
            } else {
                elementIndex += 1
                stop = doTraverseCombinations()
                elementIndex -= 1
            }

            combination = combination[:len(combination) - 1]

            if stop {
                return true
            }
        }

        if combinationLength - len(combination) < len(elements) - elementIndex {
            elementIndex += 1
            stop := doTraverseCombinations()
            elementIndex -= 1

            if stop {
                return true
            }
        }

        return false
    }

    return doTraverseCombinations()
}
