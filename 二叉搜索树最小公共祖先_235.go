func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	return helper(root, p.Val, q.Val)
}

func helper(root *TreeNode, p, q int) *TreeNode {
	r := root.Val
	if r > p && r > q {
		return helper(root.Left, p, q)
	} else if r < p && r < q {
		return helper(root.Right, p, q)
	}
	return root
}
