class Solution {
public:
    ListNode* deleteNode(ListNode* head, int val) {
        if (head == nullptr)
            return head;
        ListNode* dummy = new ListNode(0);
        dummy->next = head;
        ListNode* curr = dummy;
        while (curr->next)
        {
            if (curr->next->val == val)
            {
                curr->next = curr->next->next;
                break;
            }
            curr = curr->next;
        }
        ListNode* retNode = dummy->next;
        delete dummy;
        return retNode;
    }
};