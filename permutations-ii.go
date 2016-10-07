import (
    "fmt"
)


func permuteUnique(nums []int) [][]int {
    r := [][]int{}
    s := map[string]bool{}

    traversePermutations(len(nums), len(nums), func (a []int) bool {
        p := []int{}

        for _, i := range a {
            p = append(p, nums[i])
        }

        k := fmt.Sprintf("%v", p)

        if !s[k] {
            s[k] = true
            r = append(r, p)
        }

        return false
    })

    return r
}


func traversePermutations(elementCount int, permutationLength int, callback func ([]int) bool) bool {
    if permutationLength > elementCount {
        return false
    }

    elementIndexes := []int{}

    for i := 0; i < elementCount; i++ {
        elementIndexes = append(elementIndexes, i)
    }

    var doTraversePermutations func (int) bool

    doTraversePermutations = func (i int) bool {
        if i == permutationLength {
            return callback(elementIndexes[:permutationLength])
        }

        for j := i; j < len(elementIndexes); j++ {
            swap(&elementIndexes[i], &elementIndexes[j])
            stop := doTraversePermutations(i + 1)
            swap(&elementIndexes[i], &elementIndexes[j])

            if stop {
                return true
            }
        }

        return false
    }

    return doTraversePermutations(0)
}


func swap(x *int, y *int) {
    z := *x
    *x = *y
    *y = z
}
