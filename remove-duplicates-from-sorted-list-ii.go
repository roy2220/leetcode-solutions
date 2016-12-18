/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
    pp := &head

    for *pp != nil && (*pp).Next != nil {
        if (*pp).Val == (*pp).Next.Val {
            val := (*pp).Val
            var p *ListNode

            for p = (*pp).Next.Next; p != nil; p = p.Next {
                if p.Val != val {
                    break
                }
            }

            *pp = p
        } else {
            pp = &(*pp).Next
        }
    }

    return head
}
