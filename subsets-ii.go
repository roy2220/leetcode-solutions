import "sort"

func subsetsWithDup(nums []int) [][]int {
    sort.Ints(nums)
    c := []int{}
    r := [][]int{}
    var dfs func (i int)

    dfs = func (i int) {
        c2 := make([]int, len(c))
        copy(c2, c)
        r = append(r, c2)

        for j := i; j < len(nums); j++ {
            if j >= 1 && nums[j] == nums[j - 1] && j - 1 != i - 1 {
                continue
            }
            
            c = append(c, nums[j])
            dfs(j + 1)
            c = c[:len(c) - 1]
        }
    }

    dfs(0)
    return r
}

/*
func subsetsWithDup(nums []int) [][]int {
    sort.Ints(nums)
    r := [][]int{}
    s := []int{}
    var dfs func (i int)

    dfs = func (i int) {
        if i == len(nums) {
            s2 := make([]int, len(s))

            for i, j := range s {
                s2[i] = nums[j]
            }

            r = append(r, s2)
            return
        }

        dfs(i + 1)
        
        if !((i >= 1 && nums[i] == nums[i - 1]) && (len(s) < 1 || s[len(s) - 1] != i - 1)) {
            s = append(s, i)
            dfs(i + 1)
            s = s[:len(s) - 1]
        }
    }

    dfs(0)
    return r
}
*/
