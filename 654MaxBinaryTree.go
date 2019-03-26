/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func constructMaximumBinaryTree(nums []int) *TreeNode {

	if len(nums) == 0 {
		return nil
	}

	// find the maximum element in the nums array
	max := nums[0]
	maxIndex := 0
	for i, e := range nums {
		if max == 0 || max < e {
			max = e
			maxIndex = i
		}
	}

	// recur on left slice
	var root TreeNode
	root.Val = max
	root.Left = constructMaximumBinaryTree(nums[:maxIndex])
	// recur on right slice
	root.Right = constructMaximumBinaryTree(nums[maxIndex+1:])

	return &root

}