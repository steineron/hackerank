package datastructures

func InOrderWalk(node Node, onVisit func(n Node)) {
	if node != nil {

		InOrderWalk(node.Left(), onVisit)

		if onVisit != nil {
			onVisit(node)
		}

		InOrderWalk(node.Right(), onVisit)

	}
}

type BST struct {
	Tree
}

func (bst *BST) Search(key int) Node {
	if bst == nil {
		return nil
	}
	return search(bst.root, key)
}

func (bst *BST) Min() Node {
	if bst == nil {
		return nil
	}
	var n Node
	for n = bst.root; n.Left() != nil; n = n.Left() {
	}
	return n
}

func (bst *BST) Max() Node {
	if bst == nil {
		return nil
	}
	var n Node
	for n = bst.root; n.Right() != nil; n = n.Right() {
	}
	return n
}



func search(n Node, key int) Node {
	if n == nil || n.Key() == key {
		return n
	}

	if n.Key() < key {
		return search(n.Right(), key)
	} else {
		return search(n.Left(), key)
	}
}
