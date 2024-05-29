package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	set := make(map[*ListNode]struct{})
	cur := head
	for cur != nil && cur.Next != nil {
		if _, ok := set[cur]; ok {
			return true
		}
		set[cur] = struct{}{}
		cur = cur.Next
	}
	return false
}

func hasCycle(head *ListNode) bool {
	fast := head
	slow := head

	for slow != nil && fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if slow == fast {
			return true
		}
	}
	return false
}

