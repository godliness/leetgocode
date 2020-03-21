func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return helper(root.Left, root.Right)
}

func helper(root1, root2 *TreeNode) bool {
	// both nil
	if root1 == nil && root2 == nil {
		return true
	}

	// any nil
	if root1 == nil || root2 == nil {
		return false
	}

	return root1.Val == root2.Val && helper(root1.Right, root2.Left) && helper(root1.Left, root2.Right)
}
