/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    var l3 *ListNode
    k := 0
    prevNext := &l3

    for l1 != nil || l2 != nil || k != 0  {
        if l1 == nil {
            l1 = &ListNode{Val: 0, Next: nil}
        }

        if l2 == nil {
            l2 = &ListNode{Val: 0, Next: nil}
        }

        n3 := add(l1, l2, &k)
        *prevNext = n3
        prevNext = &n3.Next
        l1 = l1.Next
        l2 = l2.Next
    }

    *prevNext = nil
    return l3
}


func add(n1 *ListNode, n2 *ListNode, k *int) *ListNode {
    s := n1.Val + n2.Val + *k
    *k = s / 10
    n3 := ListNode{Val: s % 10}
    return &n3
}
