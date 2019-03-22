func maxDepth(root *TreeNode) int {
	// make sure tree exists
	if root == nil {
		return 0
	} else {
		// compare max depth of left and right subtree
		left := maxDepth(root.Left)
		right := maxDepth(root.Right)
		if left > right {
			return left + 1
		} else {
			return right + 1
		}
	}
}
