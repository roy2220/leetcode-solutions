/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
    ppLow := &head
    pp := &head

    for *pp != nil {
        if (*pp).Val < x {
            if pp == ppLow {
                pp = &(*pp).Next
                ppLow = pp
            } else {
                p := *pp
                *pp = (*pp).Next
                p.Next = *ppLow
                *ppLow = p
                ppLow = &p.Next
            }
        } else {
            pp = &(*pp).Next
        }
    }

    return head
}
