/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode() : val(0), left(nullptr), right(nullptr) {}
 *     TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
 *     TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left), right(right) {}
 * };
 */
class Solution {
public:
    bool helper(TreeNode* node, TreeNode* lower, TreeNode* upper) {
        if (node == nullptr) {
            return true;
        }
        if (lower != nullptr && node->val <= lower->val) {
            return false;
        }
        
        if (upper != nullptr && node->val >= upper->val) {
            return false;
        }
        return helper(node->left, lower, node) && helper(node->right, node, upper);
    }

    bool isValidBST(TreeNode* root) {
        return helper(root, nullptr, nullptr);
    }

};