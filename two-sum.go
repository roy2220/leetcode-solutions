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
