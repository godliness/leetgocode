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
class Solution {
public:
    ListNode* swapPairs(ListNode* head) {
        if(head == NULL )
            return NULL;
        if(head->next == NULL )
            return head;
        ListNode* prev = head;
        while(head && head->next){
            swap(head->val , head->next->val);
            head= head->next->next;
        }
        return prev;
      
    }
};

class Solution {
public:
    ListNode* swapPairs(ListNode* head) {
        if (head == nullptr || head->next == nullptr) return head; //基准条件
        ListNode *next = head->next;
        head->next = swapPairs(next->next); //递归处理后一部分的链表
        next->next = head;
        return next;
    }
};

class Solution {
public:
    ListNode* swapPairs(ListNode* head) {
        if (head == nullptr || head->next == nullptr) {
            return head;
        }
        
        ListNode* dummy = new ListNode(0);
        dummy->next = head;
        ListNode* curr = dummy;
        
        while (curr->next != nullptr && curr->next->next != nullptr) {
            ListNode* first = curr->next;
            ListNode* second = curr->next->next;
            curr->next = second;
            first->next = second->next;
            second->next = first;
            curr = curr->next->next;
        }
        
        return dummy->next;
    }
};
