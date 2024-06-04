class Solution {
public:
    ListNode* removeElements(ListNode* head, int val) {
        while (head && head->val == val) {
            head = head->next;
        }

        if (!head) {
            return head;
        }


        ListNode* pre = head;
        while (pre->next) {
            if (pre->next->val == val) {
                pre->next = pre->next->next;
            } else {
                pre = pre->next;
            }
        }
        return head;
    }
};
