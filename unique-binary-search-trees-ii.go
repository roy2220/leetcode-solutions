/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func generateTrees(n int) []*TreeNode {
    if n == 0 {
        return []*TreeNode{}
    }

    return build(1, n)
}

func build(i int, j int) []*TreeNode {
    if i == j {
        return []*TreeNode{&TreeNode{Val: i, Left: nil, Right: nil}}
    } else if i > j {
        return []*TreeNode{nil}
    }

    ts := []*TreeNode{}

    for k := i; k <= j; k++ {
        ts1 := build(i, k - 1)
        ts2 := build(k + 1, j)

        for _, l := range ts1 {
            for _, r := range ts2 {
                ts = append(ts, &TreeNode{Val: k, Left: l, Right: r})
            }
        }
    }

    return ts
}
