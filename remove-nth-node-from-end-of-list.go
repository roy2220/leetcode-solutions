func removeNthFromEnd(head *ListNode, n int) *ListNode {
    l := 0

    for node := head; node != nil; node = node.Next {
        l++
    }

    prevNext := &head

    for i := 0; i < l - n; i++ {
        prevNext = &(*prevNext).Next
    }

    *prevNext = (*prevNext).Next
    return head
}
