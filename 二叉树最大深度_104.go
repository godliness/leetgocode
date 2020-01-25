func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    
    return 1 + max(maxDepth(root.Right), maxDepth(root.Left))
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}
