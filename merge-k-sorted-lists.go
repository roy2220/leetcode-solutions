/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

import (
    "container/heap"
)


type Heap []*ListNode


func (h Heap) Len() int           { return len(h) }
func (h Heap) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h Heap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }


func (h *Heap) Push(x interface{}) {
    *h = append(*h, x.(*ListNode))
}


func (h *Heap) Pop() interface{} {
    l := len(*h)
    x := (*h)[l - 1]
    *h = (*h)[0:l - 1]
    return x
}


func mergeKLists(lists []*ListNode) *ListNode {
    var r *ListNode = nil
    prevNext := &r
    h := &Heap{}
    heap.Init(h)

    for _, l := range lists {
        if l != nil {
            heap.Push(h, l)
            l = l.Next
        }
    }

    for h.Len() >= 1 {
        n := (*h)[0]

        if n.Next == nil {
            heap.Pop(h)
        } else {
            (*h)[0] = n.Next
            heap.Fix(h, 0)
        }

        *prevNext = n
        prevNext = &n.Next
    }

    return r
}
