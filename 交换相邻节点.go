package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) (r *ListNode) {
	if head == nil || head.Next == nil {
		return head
	}
	r, head.Next, head.Next.Next = head.Next, swapPairs(head.Next.Next), head
	return r
}
