func hasPathSum(root *TreeNode, sum int) bool {
    if root == nil {
        return false
    }
    if root.Left == nil && root.Right == nil {
        return root.Val == sum
    }
    if root.Left != nil {
        if hasPathSum(root.Left, sum - root.Val) {
            return true
        }
    }
    if root.Right != nil {
        if hasPathSum(root.Right, sum - root.Val) {
            return true
        }
    }
    return false
}
