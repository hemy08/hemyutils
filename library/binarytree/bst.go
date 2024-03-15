package binarytree

// InsertBST 插入一个值到二叉搜索树中
func (node *BTreeNode) InsertBST(value int) {
	if node == nil {
		return
	}

	if value < node.Val {
		if node.Left == nil {
			node.Left = NewBTreeNode(value)
		} else {
			node.Left.InsertBST(value)
		}
	} else {
		if node.Right == nil {
			node.Right = NewBTreeNode(value)
		} else {
			node.Right.InsertBST(value)
		}
	}
}

// NewBST 创建一个新的二叉搜索树
func NewBST() *BTreeNode {
	return nil
}
