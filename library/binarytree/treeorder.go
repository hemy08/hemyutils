package binarytree

func PreOrderBinaryTree(root *BTreeNode, dep int, maxDep *int) {
	if dep > *maxDep {
		*maxDep = dep
	}
	if root == nil {
		return
	}
	PreOrderBinaryTree(root.Left, dep+1, maxDep)
	PreOrderBinaryTree(root.Right, dep+1, maxDep)
}

// PreOrderTraversal 前序遍历： 根-> 左子树 -> 右子树
func PreOrderTraversal(root *BTreeNode, res *[]int) {
	if root == nil {
		return
	}
	*res = append(*res, root.Val)
	PreOrderTraversal(root.Left, res)
	PreOrderTraversal(root.Right, res)
}

// IntermediateOrderTraversal 中序： 左子树-> 根 -> 右子树
func IntermediateOrderTraversal(root *BTreeNode, res *[]int) {
	if root == nil {
		return
	}
	IntermediateOrderTraversal(root.Left, res)
	*res = append(*res, root.Val)
	IntermediateOrderTraversal(root.Right, res)
}

// PostOrderTraversal 后序： 左子树-> 右子树 ->  根
func PostOrderTraversal(root *BTreeNode, res *[]int) {
	if root == nil {
		return
	}
	PostOrderTraversal(root.Left, res)
	PostOrderTraversal(root.Right, res)
	*res = append(*res, root.Val)
}

func helperLevelOrderTraversal(nodes []*BTreeNode, data *[][]int) {
	if nodes == nil {
		return
	}

	var newNodes []*BTreeNode
	var result []int
	for _, v := range nodes {
		if v == nil {
			return
		}

		if v.Left != nil {
			newNodes = append(newNodes, v.Left)
		}

		if v.Right != nil {
			newNodes = append(newNodes, v.Right)
		}
		result = append(result, v.Val)
	}
	*data = append(*data, result)
	helperLevelOrderTraversal(newNodes, data)
}

// LevelOrderTraversal 层序遍历
func LevelOrderTraversal(root *BTreeNode, res *[][]int) {
	nodes := []*BTreeNode{root}
	helperLevelOrderTraversal(nodes, res)
}
