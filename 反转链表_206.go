package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		head.Next, prev, head = prev, head, head.Next
	}
	return prev
}

func reverseList(head *ListNode) *ListNode {
    var prev *ListNode
    for head != nil {
        nextTemp := head.Next
        head.Next = prev
        prev = head
        head = nextTemp
    }
    return prev
}
