package datastructures

type Node interface {
	Key() int
	Left() Node
	Right() Node
}

type treeNode struct {
	key         int
	left, right *treeNode
}

func (t *treeNode) Key() int {
	return t.key
}

func (t *treeNode) Left() Node {
	return t.left
}

func (t *treeNode) Right() Node {
	return t.right
}

type Tree struct {
	root *treeNode
}
