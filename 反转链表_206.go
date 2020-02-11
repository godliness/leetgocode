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
        head.Next = prev // head.Next 指针要指向前面的节点，而且prev变量的意义是接收head的，而head要被赋值成head.Next，以供下一次循环使用。
        prev = head
        head = nextTemp // 典型额 head = head.Next 遍历链表
    }
    return prev
}
