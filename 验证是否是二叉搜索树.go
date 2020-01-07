func isValidBST(root *TreeNode) bool {
    return helper(root, nil, nil)
}

func helper(root, l, r *TreeNode) bool {
    if root == nil { return true }
    left, right := true, true
    if l != nil { left = l.Val < root.Val }
    if r != nil { right = root.Val < r.Val }
    return left && right && helper(root.Left, l, root) && helper(root.Right, root, r)
}
