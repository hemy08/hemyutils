package binarytree

import (
	"golang.org/x/exp/errors/fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_BinaryTree_MaxDepth_Case1(t *testing.T) {
	root := &BTreeNode{
		Val: 3,
		Left: &BTreeNode{
			Val:   9,
			Left:  nil,
			Right: nil,
		},
		Right: &BTreeNode{
			Val:   20,
			Left:  &BTreeNode{Val: 15, Left: nil, Right: nil},
			Right: &BTreeNode{Val: 7, Left: nil, Right: nil},
		},
	}

	res := MaxDepth(root)
	assert.Equal(t, 3, res)
	println("res :", res)
}

func Test_BinaryTree_MaxDepth_Case2(t *testing.T) {
	root := &BTreeNode{
		Val:   1,
		Left:  nil,
		Right: &BTreeNode{Val: 2, Left: nil, Right: nil},
	}

	res := MaxDepth(root)
	assert.Equal(t, 2, res)
	println("res :", res)
}

func Test_BinaryTree_PreOrder(t *testing.T) {
	btree := NewBinaryTree(NewBTreeNode(5))
	values := []int{3, 7, 2, 4, 6, 8}
	for _, value := range values {
		btree.Root.InsertBST(value)
	}

	_, _ = fmt.Println("res", btree.PreOrderTraversal())
	_, _ = fmt.Println("res", btree.IntermediateOrderTraversal())
	_, _ = fmt.Println("res", btree.PostOrderTraversal())
}

func Test_BinaryTree_IsSymmetric_Case1(t *testing.T) {
	root := &BTreeNode{
		Val: 1,
		Left: &BTreeNode{
			Val:   2,
			Left:  nil,
			Right: &BTreeNode{Val: 3, Left: nil, Right: nil}},
		Right: &BTreeNode{
			Val:   2,
			Left:  nil,
			Right: &BTreeNode{Val: 3, Left: nil, Right: nil}},
	}

	res := IsSymmetric(root)
	assert.Equal(t, false, res)
	println("res :", res)
}

func Test_BinaryTree_IsSymmetric_Case2(t *testing.T) {
	root := &BTreeNode{
		Val: 1,
		Left: &BTreeNode{
			Val:   0,
			Left:  nil,
			Right: nil},
		Right: nil,
	}

	res := IsSymmetric(root)
	assert.Equal(t, false, res)
	println("res :", res)
}

func Test_BinaryTree_IsSymmetric_Case3(t *testing.T) {
	root := &BTreeNode{
		Val: 1,
		Left: &BTreeNode{
			Val: 2,
			Left: &BTreeNode{
				Val:   2,
				Left:  nil,
				Right: nil},
			Right: nil},
		Right: &BTreeNode{
			Val: 2,
			Left: &BTreeNode{
				Val:   2,
				Left:  nil,
				Right: nil},
			Right: nil},
	}

	res := IsSymmetric(root)
	assert.Equal(t, false, res)
	println("res :", res)
}

func Test_BinaryTree_IsSymmetric_Case4(t *testing.T) {
	root := &BTreeNode{
		Val: 5,
		Left: &BTreeNode{
			Val:  4,
			Left: nil,
			Right: &BTreeNode{
				Val: 1,
				Left: &BTreeNode{
					Val:   2,
					Left:  nil,
					Right: nil},
				Right: nil}},
		Right: &BTreeNode{
			Val:  1,
			Left: nil,
			Right: &BTreeNode{
				Val: 4,
				Left: &BTreeNode{
					Val:   2,
					Left:  nil,
					Right: nil},
				Right: nil}},
	}

	res := IsSymmetric(root)
	assert.Equal(t, false, res)
	println("res :", res)
}
