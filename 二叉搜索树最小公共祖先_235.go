func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	return helper(root, p.Val, q.Val)
}

// 二叉树特点，右子树永远大于根节点，左子树永远小于根节点
func helper(root *TreeNode, p, q int) *TreeNode {
	r := root.Val //从上往下找
	if r > p && r > q { // 根节点如果同时大于p, q 那么说明公共祖先应该在左子树，说明找反了
		return helper(root.Left, p, q) 
	} else if r < p && r < q { // 根节点如果同时小于p, q 那么说明公共祖先应该在右子树，说明找反了
		return helper(root.Right, p, q)
	}
	return root // root正好 > p && < q || <p && > q 说明找到了 p q 正好在root的两边
}
