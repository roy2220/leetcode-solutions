/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
    pp1 := &head

    for {
        node1 := *pp1

        for i := 0; i < k; i++ {
            if node1 == nil {
                return head
            }

            node1 = node1.Next
        }

        node2 := *pp1
        node3 := node2
        node2 = node2.Next
        node3.Next = node1
        pp2 := &node3.Next

        for node2 != node1 {
            node3 = node2
            node2 = node2.Next
            node3.Next = *pp1
            *pp1 = node3
        }

        pp1 = pp2
    }

    return head
}
