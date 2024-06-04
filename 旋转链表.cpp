/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode() : val(0), next(nullptr) {}
 *     ListNode(int x) : val(x), next(nullptr) {}
 *     ListNode(int x, ListNode *next) : val(x), next(next) {}
 * };
 */
// 链表首尾相连，走k步，第k个结点下一个指向nullptr.
class Solution {
public:
    ListNode* rotateRight(ListNode* head, int k) {
         if (k == 0 || head == nullptr || head->next == nullptr) {
            return head;
        }
        int n = 1;
        ListNode* curr = head;
        while (curr->next != nullptr) {
            curr = curr->next;
            n++;
        }
        int add = n - k % n;
        if (add == n) {
            return head;
        }
        curr->next = head;
        while (add--) {
            curr = curr->next;
        }
        ListNode* ret = curr->next;
        curr->next = nullptr;
        return ret;
    }
};