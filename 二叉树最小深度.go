func minDepth(root *TreeNode) int {
    if root == nil {
	return 0 // 到头了
    }
    left, right := minDepth(root.Left), minDepth(root.Right) 
    if left == 0 || right == 0 { // 左面和右面不断的往下递归 一旦发现有子节点到头了，立马返回当前层+1就是结果
	return 1 + left + right
    }
    return 1 + min(left, right) // 因为左右子树都会到头 取到头后最小的值
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}
