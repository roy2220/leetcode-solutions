/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
    var prevNext = &head

    for {
        n1 := *prevNext

        if n1 == nil {
            break
        }

        n2 := n1.Next

        if n2 == nil {
            break
        }

        n1.Next = n2.Next
        n2.Next = n1
        *prevNext = n2
        prevNext = &n1.Next
    }

    return head
}
