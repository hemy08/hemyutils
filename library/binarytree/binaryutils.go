package binarytree

// GetBTreeDepth 二叉树深度
func GetBTreeDepth(root *BTreeNode) int {
	if root == nil {
		return 0
	}
	var maxDegree = 0
	if GetBTreeDepth(root.Left) > GetBTreeDepth(root.Right) {
		maxDegree = GetBTreeDepth(root.Left)
	} else {
		maxDegree = GetBTreeDepth(root.Right)
	}
	return maxDegree + 1
}

// GetKthNodeNum 求 K 层节点个数
func GetKthNodeNum(root *BTreeNode, k int) int {
	if root == nil {
		return 0
	}
	if k == 1 {
		return 1
	}
	return GetKthNodeNum(root.Left, k-1) + GetKthNodeNum(root.Right, k-1)
}

// GetLeafNums 求叶子节点个数
func GetLeafNums(root *BTreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	return GetLeafNums(root.Left) + GetLeafNums(root.Right)
}

// MaxDepth 最大深度
func MaxDepth(root *BTreeNode) int {
	depth := 0
	PreOrderBinaryTree(root, 0, &depth)
	return depth
}

// MaxDepthV2 最大深度
func MaxDepthV2(root *BTreeNode) int {
	if root == nil {
		return 0
	}

	left := MaxDepthV2(root.Left)
	right := MaxDepthV2(root.Right)

	if left > right {
		return left + 1
	}

	return right + 1
}
