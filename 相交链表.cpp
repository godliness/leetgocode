class Solution {
public:
    ListNode *getIntersectionNode(ListNode *headA, ListNode *headB) {
       if (headA == nullptr || headB == nullptr) {
            return nullptr;
        }
        ListNode *pA = headA, *pB = headB;
        while (pA != pB) {
            pA = pA == nullptr ? headB : pA->next;
            pB = pB == nullptr ? headA : pB->next;
        }
        return pA;
    }
};

// 链表存在交点的情况：若两链表长度相同，那么两个指针会在移动的过程中相遇;但是若两个链表长度不相同，则它们第一次遍历无法相遇。继续规定第一个指针遍历完第一个链表后遍历第二个链表，第二个指针遍历完第二个链表后遍历第一个链表，如此肯定能保证两个指针到达交点的路程相同，则它们会相遇。

// 链表无交点的情况：若两个链表长度相同，则两个指针遍历完各自的链表，最后同时指向 nullptr 可以判断不存在交点。若两个链表长度不相同，则两个指针遍历完各自的链表遍历另一个链表，最后也会同时指向 nullptr 可以判断不存在交点。