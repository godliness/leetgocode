package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) (r *ListNode) {
	if head == nil || head.Next == nil {
		return head
	}
	var r *ListNode
	r, head.Next, head.Next.Next = head.Next, swapPairs(head.Next.Next), head
	return r
}

func swapPairs2(head *ListNode) *ListNode {
    newH := &ListNode{}
    newH.Next = head
    cur := head
    prev := newH
    
    for cur != nil && cur.Next != nil {
        prev.Next = cur.Next
        temp := cur.Next.Next
        cur.Next.Next = cur
        cur.Next = temp
        prev = cur
        cur = temp
    }
    
    return newH.Next
}
