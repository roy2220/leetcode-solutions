func twoSum(nums []int, target int) []int {
    result := []int{}

    traverseCombinations(nums, 2, func (combination []int) bool {
        i, j := combination[0], combination[1]

        if nums[i] + nums[j] == target {
            result = combination
            return true
        } else {
            return false
        }
    })

    return result
}


func traverseCombinations(elements []int, combinationLength int, callback func ([]int) bool) bool {
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
