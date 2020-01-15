func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    
    right := maxDepth(root.Right)
    left := maxDepth(root.Left)
    
    if right > left {
        return right + 1
    } 
    
    return left + 1
}
