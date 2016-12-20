/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
    r := []int{}
    var dfs func (*TreeNode)

    dfs = func (p *TreeNode) {
        if p == nil {
            return
        }

        dfs(p.Left)
        r = append(r, p.Val)
        dfs(p.Right)
    }

    dfs(root)
    return r
}
