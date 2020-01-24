func minDepth(root *TreeNode) int {
    if root == nil {
	return 0
    }
    left, right := minDepth(root.Left), minDepth(root.Right)
    if left == 0 || right == 0 {
	return 1 + left + right
    }
    return 1 + min(left, right)
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}
