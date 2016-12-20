/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
    var lastVal int
    lastValSet := false
    var dfs func (*TreeNode) bool

    dfs = func (p *TreeNode) bool {
        if p == nil {
            return true
        }

        if !dfs(p.Left) {
            return false
        }

        if lastValSet && p.Val <= lastVal {
            return false
        }

        lastVal = p.Val
        lastValSet = true

        if !dfs(p.Right) {
            return false
        }

        return true
    }

    return dfs(root)
}
