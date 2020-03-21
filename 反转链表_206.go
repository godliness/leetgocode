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
        nextTemp := head.Next // 先预留head.Next以后用
        head.Next = prev // head.Next 指针要指向前面的节点，而且prev变量的意义是接收head的，而head要被赋值成head.Next，以供下一次循环使用。
        prev = head
        head = nextTemp // 典型的 head = head.Next 遍历链表
    }
    return prev
}

func reverseList(head *ListNode) *ListNode {
     if head == nil || head.Next == nil {
         return head
     }
     reverse := reverseList(head.Next)
     head.Next.Next = head
     head.Next = nil
     return reverse
}
