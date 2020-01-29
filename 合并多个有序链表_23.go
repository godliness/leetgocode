func mergeKLists(lists []*ListNode) *ListNode {
    if len(lists) == 0 {
        return nil
    }
    var sliceList []int
    for _, v := range lists {
        for v != nil {
            sliceList = append(sliceList, v.Val)
            v = v.Next
        }
    }
    if sliceList == nil {
        return nil
    }
    sort.Ints(sliceList)
    head := &ListNode{}
    cur := head
    for _, v := range sliceList {
        cur.Next = &ListNode{Val:v}
        cur = cur.Next
    }
    return head.Next
}
