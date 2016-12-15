import (
    "sort"
    "fmt"
)

func getPermutation(n int, k int) string {
    a := []int{}
    s := 1

    for m := n - 1; m >= 2; m-- {
        s *= m
    }

    if n >= 1 {
        m := n - 1

        for {
            i := ((k - 1) / s) + 1
            b := make([]int, len(a))
            copy(b, a)
            sort.Ints(b)

            for _, v := range b {
                if i < v {
                    break
                }

                i++
            }

            a = append(a, i)

            if m == 0 {
                break
            }

            k = ((k - 1) % s) + 1
            s /= m
            m--
        }
    }

    r := ""

    for _, v := range a {
        r += fmt.Sprintf("%v", v);
    }

    return r
}
