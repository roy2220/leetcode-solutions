iimport (
    "sort"
    "fmt"
)


func threeSum(nums []int) [][]int {
    result := [][]int{}
    sort.Ints(nums)
    s := map[string]bool{}

    traverseCombinations(nums, 3, func (combination []int) bool {
        i, j, k := combination[0], combination[1], combination[2]
        x, y, z := nums[i], nums[j], nums[k]

        if x + y + z == 0 {
            key := fmt.Sprintf("%d,%d,%d", x,y,z)

            if !s[key] {
                s[key] = true
                result = append(result, []int{x, y, z})
            }
        }

        return false
    })

    return result
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
