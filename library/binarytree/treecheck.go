package binarytree

import "math"

func SubTreeValidBST(root *BTreeNode, min, max int) bool {
	if root == nil {
		return true
	}

	if root.Val <= min || root.Val >= max {
		return false
	}

	left := SubTreeValidBST(root.Left, min, root.Val)
	right := SubTreeValidBST(root.Right, root.Val, max)

	return left && right
}

// IsValidBST 二叉搜索树
func IsValidBST(root *BTreeNode) bool {
	return SubTreeValidBST(root, math.MinInt64, math.MaxInt64)
}

// IsSymmetric 对称二叉树
func IsSymmetric(root *BTreeNode) bool {
	var left, right []int

	// root左侧进行前序遍历
	PreOrderTraversal(root.Left, &left)
	// root右侧进行后序遍历
	PostOrderTraversal(root.Right, &right)
	if len(left) != len(right) {
		return false
	}

	for i := 0; i < len(left); i++ {
		if left[i] != right[i] {
			return false
		}
	}

	return true
}

// IsBalancedTree 平衡二叉树
func IsBalancedTree(root *BTreeNode) bool {
	if root == nil {
		return true
	}
	lde := GetBTreeDepth(root.Left)
	rde := GetBTreeDepth(root.Right)
	flag := false
	if (math.Abs(float64(lde - rde))) <= 1 {
		flag = true
	} else {
		flag = false
	}
	return flag && IsBalancedTree(root.Left) && IsBalancedTree(root.Right)
}
