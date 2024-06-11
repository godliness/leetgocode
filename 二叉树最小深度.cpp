class Solution {
public:
    int minDepth(TreeNode* root) {
        if(root==nullptr){
            // 一个孩子节点都没有
            return 0;
        }
        if(root->left&&root->right){
            // 如果这个节点的左右孩子都存在
            return min(minDepth(root->left),minDepth(root->right))+1;
        }
        // 如果当前节点只有一个孩子，只需一条道走到黑
        return root->right==nullptr?minDepth(root->left)+1:minDepth(root->right)+1;
    }
};