/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    var l3 *ListNode = nil
    prevNext := &l3

    for l1 != nil && l2 != nil {
        if l1.Val < l2.Val {
            *prevNext = l1
            prevNext = &l1.Next
            l1 = l1.Next
        } else {
            *prevNext = l2
            prevNext = &l2.Next
            l2 = l2.Next
        }
    }

    if l2 != nil {
        *prevNext = l2
    } else if l1 != nil {
        *prevNext = l1
    }

    return l3
}
