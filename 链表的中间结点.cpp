class Solution {
public:
    ListNode* middleNode(ListNode* head) {
        vector<ListNode*> nodes;
        while(head)
        {
            nodes.emplace_back(head);
            head = head->next;
        }

        for (int i =0, j= nodes.size()-1; i <= j; i++, j--)
        {
            if (i == j)
                return nodes[i];
            else if (j - i == 1)
                return nodes[j];
        }
        return nullptr;
    }
};