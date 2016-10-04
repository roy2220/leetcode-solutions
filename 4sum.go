import (
    "sort"
    "fmt"
)


func fourSum(nums []int, target int) [][]int {
    result := [][]int{}
    m := map[int][]int{}

    for idx, num := range nums {
        if a, ok := m[num]; ok {
            m[num] = append(a, idx)
        } else {
            m[num] = []int{idx}
        }
    }

    s := map[string]bool{}

    traverseCombinations(nums, 3, func (combination []int) bool {
        i, j, k := combination[0], combination[1], combination[2]
        w, x, y := nums[i], nums[j], nums[k]
        z := target - (w + x + y)

        if a, ok := m[z]; ok {
            f := len(a)

            for _, idx := range a {
                if idx == i || idx == j || idx == k {
                    f--
                }
            }

            if f >= 1 {
                val := []int{w, x, y, z}
                sort.Ints(val)
                key := fmt.Sprintf("%v", val)

                if !s[key] {
                    s[key] = true
                    result = append(result, val)
                }
            }
        }

        return false
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
