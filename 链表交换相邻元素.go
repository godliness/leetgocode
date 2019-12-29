func swapPairs(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head;
    }
    res := head.Next
    head.Next = swapPairs(res.Next)
    res.Next = head
    return res
}

func swapPairs(head *ListNode) *ListNode {
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

