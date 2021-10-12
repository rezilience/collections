package collections

type treeNode struct {
	data   int
	parent *treeNode
	left   *treeNode
	right  *treeNode
	isRed  bool // default/false is considered as black
}

/*
A red-black tree T is a binary search tree having following five additional properties (invariants).
1. Every node in T is either red or black.
2. The root node of T is black.
3. Every nil node is black. (Nil nodes are the leaf nodes. They do not contain any keys.
	When we search for a key that is not present in the tree, we reach the nil node.)
4. If a node is red, its parent and both children are black.
	This means no two adjacent nodes on a path can be red nodes.
5. Every path from a root node to a nil node has the same number of black nodes.

Reference - https://algorithmtutor.com/Data-Structures/Tree/Red-Black-Trees/
*/
type RedBlackTree struct {
	root *treeNode
}

func (r *RedBlackTree) Insert(data int) {
	n := &treeNode{data: data}
	r.insertInBST(n)              // Ordinary BST insertion
	r.ensureRedBlackProperties(n) // Ensure all 5 properties of red-black tree hold.
}

func (r *RedBlackTree) insertInBST(n *treeNode) {
	var curr, prev *treeNode = r.root, nil

	for curr != nil {
		prev = curr
		if n.data < curr.data {
			curr = curr.left
		} else {
			curr = curr.right
		}
	}

	n.parent = prev
	if prev == nil {
		r.root = n
	} else if n.data < prev.data {
		prev.left = n
	} else {
		prev.right = n
	}
}

func (r *RedBlackTree) ensureRedBlackProperties(n *treeNode) {
	if n.parent == nil {
		// This is the root of the tree. It can't violate red-black tree properties, as long as it's black.
		return
	}
	n.isRed = true
	if n.parent.parent == nil { // If the parent is root (black), this insertion can't violate any properties as long as it's red.
		return
	}
	r.fixInsert(n)
}

// fixInsert ensures that the tree is balanced and maintains the 5 properties of a red-black tree.
func (r *RedBlackTree) fixInsert(n *treeNode) {
	// Assumes parent and grandparent are not nil.
	if n.parent == nil || n.parent.parent == nil {
		panic("this shouldn't have happened!")
	}
	for n.parent != nil && n.parent.isRed { // If parent is red, it violates property no. 4 of red-black trees.
		var u, p, g *treeNode = nil, n.parent, n.parent.parent
		if p == g.right {
			u = g.left
		} else {
			u = g.right
		}

		if u != nil && u.isRed { // 3.1
			// Flip the colors of P, U and G. That is, push blackness downstream.
			p.isRed, u.isRed, g.isRed = false, false, true
			n = g
			continue
		}

		if p == g.right && n == p.right { // 3.2.1
			// Single left rotation at g
			r.leftRotate(g)

			// Flip the colors of the new sibling and parent
			n.parent.isRed = false
			n.parent.left.isRed = true
		} else if p == g.right && n == p.left { // 3.2.2
			n = p
			r.rightRotate(p)
		} else if n == p.left { // 3.2.3
			r.rightRotate(g)

			n.parent.isRed = false
			n.parent.right.isRed = true
		} else { // 3.2.4
			n = p
			r.leftRotate(p)
		}
	}

	r.root.isRed = false
}

func (r *RedBlackTree) leftRotate(me *treeNode) {
	// It assumes that me is NOT nil, and that me.right is NOT nil.
	if me == nil || me.right == nil {
		panic("this should not have happened!")
	}

	// Rotate (demote) me to the left
	parent, rightChild := me.parent, me.right
	me.parent = rightChild
	me.right = rightChild.left

	// Fix right child's left child's parent
	if rightChild.left != nil {
		rightChild.left.parent = me
	}

	// Rotate (promote) the right child to the left
	if parent == nil {
		r.root = rightChild
	} else if parent.left == me {
		parent.left = rightChild
	} else {
		parent.right = rightChild
	}
	rightChild.left = me
	rightChild.parent = parent
}

func (r *RedBlackTree) rightRotate(me *treeNode) {
	// It assumes that me is NOT nil, and that me.left is NOT nil.
	if me == nil || me.left == nil {
		panic("this should not have happened!")
	}

	// Rotate (demote) me to the right
	parent, leftChild := me.parent, me.left
	me.parent = leftChild
	me.left = leftChild.right

	// Fix left child's right child's parent
	if leftChild.right != nil {
		leftChild.right.parent = me
	}

	// Rotate (promote) the right child to the left
	if parent == nil {
		r.root = leftChild
	} else if parent.left == me {
		parent.left = leftChild
	} else {
		parent.right = leftChild
	}
	leftChild.right = me
	leftChild.parent = parent
}
