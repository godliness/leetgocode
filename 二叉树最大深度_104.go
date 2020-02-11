func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    
    return 1 + max(maxDepth(root.Right), maxDepth(root.Left)) // 一直递归到最末尾，一次上翻+1，得到最大深度
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}
