func removeElements(head *ListNode, val int) *ListNode {
    result := &ListNode{0, head}
    tmp := result
    for tmp.Next != nil; {
        if tmp.Next.Val == val {
            tmp.Next = tmp.Next.Next
        }else {
            tmp = tmp.Next
        }
    }

	return result.Next
}
