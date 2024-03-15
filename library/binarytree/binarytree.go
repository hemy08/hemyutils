package binarytree

type BTreeNode struct {
	Val   int
	Left  *BTreeNode
	Right *BTreeNode
}

type BinaryTree struct {
	Root                               *BTreeNode
	leafNums, depth, maxDepth, kthNums int
	preOrder, midOrder, postOrder      []int
	levelOrder                         [][]int
	isBST, balance, isSymmetric        bool
}

// NewBinaryTree 新创建二叉树
func NewBinaryTree(node *BTreeNode) *BinaryTree {
	return &BinaryTree{
		Root:       node,
		preOrder:   make([]int, 0),
		midOrder:   make([]int, 0),
		postOrder:  make([]int, 0),
		levelOrder: make([][]int, 0),
	}
}

// NewBTreeNode 新创建二叉树节点
func NewBTreeNode(val int) *BTreeNode {
	return &BTreeNode{
		val, nil, nil,
	}
}

// NodeNum 计算二叉树节点个数
func (b *BinaryTree) NodeNum(Root *BTreeNode) int {
	if Root == nil {
		return 0
	} else {
		return b.NodeNum(Root.Left) + b.NodeNum(Root.Right) + 1
	}
}

// Depth 计算二叉树的深度
func (b *BinaryTree) Depth() int {
	if b.Root == nil {
		return 0
	}
	b.depth = GetBTreeDepth(b.Root)
	return b.depth
}

// MaxDepthV1 最大深度
func (b *BinaryTree) MaxDepthV1() int {
	if b.Root == nil {
		return 0
	}
	b.maxDepth = MaxDepth(b.Root)
	return b.maxDepth
}

// MaxDepthV2 最大深度
func (b *BinaryTree) MaxDepthV2() int {
	if b.Root == nil {
		return 0
	}
	b.maxDepth = MaxDepthV2(b.Root)
	return b.maxDepth
}

// LeafNums 求叶子节点个数
func (b *BinaryTree) LeafNums() int {
	b.leafNums = GetLeafNums(b.Root)
	return b.leafNums
}

// KthNodeNum 求 K 层节点个数
func (b *BinaryTree) KthNodeNum(k int) int {
	b.kthNums = GetKthNodeNum(b.Root, k)
	return b.kthNums
}

// PreOrderTraversal 前序遍历： 根-> 左子树 -> 右子树
func (b *BinaryTree) PreOrderTraversal() []int {
	if b.Root == nil {
		return nil
	}

	if len(b.preOrder) == 0 {
		PreOrderTraversal(b.Root, &b.preOrder)
	}

	return b.preOrder
}

// IntermediateOrderTraversal 中序： 左子树-> 根 -> 右子树
func (b *BinaryTree) IntermediateOrderTraversal() []int {
	if b.Root == nil {
		return nil
	}

	if len(b.midOrder) == 0 {
		IntermediateOrderTraversal(b.Root, &b.midOrder)
	}

	return b.midOrder
}

// PostOrderTraversal 后序： 左子树-> 右子树 ->  根
func (b *BinaryTree) PostOrderTraversal() []int {
	if b.Root == nil {
		return nil
	}

	if len(b.postOrder) == 0 {
		PostOrderTraversal(b.Root, &b.postOrder)
	}

	return b.postOrder
}

// LevelOrderTraversal 层序遍历
func (b *BinaryTree) LevelOrderTraversal() [][]int {
	if len(b.levelOrder) == 0 {
		LevelOrderTraversal(b.Root, &b.levelOrder)
	}

	return b.levelOrder
}

// IsBalancedTree 判断是否平衡二叉树
func (b *BinaryTree) IsBalancedTree() bool {
	b.balance = IsBalancedTree(b.Root)
	return b.balance
}

// IsValidBST 判断是否二叉搜索树
func (b *BinaryTree) IsValidBST() bool {
	b.isBST = IsValidBST(b.Root)
	return b.isBST
}

// IsSymmetric 判断是否对称二叉树
func (b *BinaryTree) IsSymmetric() bool {
	b.isSymmetric = IsSymmetric(b.Root)
	return b.isSymmetric
}
