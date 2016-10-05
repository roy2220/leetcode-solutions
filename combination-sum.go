import (
    "sort"
)


func combinationSum(candidates []int, target int) [][]int {
    r := [][]int{}
    m := []int{}

    var f func (int, int)

    f = func (i int, t int) {
        if t == 0 {
            n := make([]int, len(m))
            copy(n, m)
            r = append(r, n)
        }

        pc := -1

        for j := i; j < len(candidates); j++ {
            c := candidates[j]

            if c > t {
                break
            }

            if c != pc {
                m = append(m, c)
                f(j, t - c)
                m = m[:len(m) - 1]
                pc = c
            }
        }
    }

    sort.Ints(candidates)
    f(0, target)
    return r
}
