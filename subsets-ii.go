import "sort"

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
