func subsets(nums []int) [][]int {
    r := [][]int{}
    s := []int{}
    var dfs func (i int)

    dfs = func (i int) {
        if i == len(nums) {
            s2 := make([]int, len(s))
            copy(s2, s)
            r = append(r, s2)
            return
        }

        dfs(i + 1)
        s = append(s, nums[i])
        dfs(i + 1)
        s = s[:len(s) - 1]
    }

    dfs(0)
    return r
}
