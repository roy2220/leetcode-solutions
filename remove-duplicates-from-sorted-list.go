/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }

    pp := &head

    for (*pp).Next != nil {
        if (*pp).Val == (*pp).Next.Val {
            *pp = (*pp).Next
        } else {
            pp = &(*pp).Next
        }
    }

    return head
}
