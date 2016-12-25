import "sort"

func permute(nums []int) [][]int {
    sort.Ints(nums)
    r := [][]int{}
    var dfs func (i int)

    dfs = func (i int) {
        if i == len(nums) {
            c := make([]int, len(nums))
            copy(c, nums)
            r = append(r, c)
            return
        }

        for j := i; j < len(nums); j++ {
            nums[i], nums[j] = nums[j], nums[i]
            dfs(i + 1)
            nums[i], nums[j] = nums[j], nums[i]
        }
    }

    dfs(0)
    return r
}
