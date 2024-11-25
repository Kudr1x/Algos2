package trees

import (
	"fmt"
)

type Color int

const (
	RED   = Color(0)
	BLACK = Color(1)
)

type rbtNode struct {
	key    int64
	value  string
	color  Color
	parent *rbtNode
	left   *rbtNode
	right  *rbtNode
}

type RedAndBlackTree struct {
	root *rbtNode
	size int
}

func NewTree() *RedAndBlackTree {
	return &RedAndBlackTree{}
}

func (t *RedAndBlackTree) Root() *rbtNode {
	return t.root
}

func (t *RedAndBlackTree) Insert(key int64, value string) {
	x := newNode(key, value)
	t.insert(x)
}

func (t *RedAndBlackTree) insert(item *rbtNode) {
	var y *rbtNode
	x := t.root

	for x != nil {
		y = x
		if item.key < x.key {
			x = x.left
		} else if item.key > x.key {
			x = x.right
		} else {
			return
		}
	}

	t.size++
	item.parent = y
	item.color = RED

	if y == nil {
		item.color = BLACK
		t.root = item
		return
	} else if item.key < y.key {
		y.left = item
	} else {
		y.right = item
	}

	t.insertRepairNode(item)
}

func (t *RedAndBlackTree) insertRepairNode(x *rbtNode) {
	var y *rbtNode
	for x != t.root && x.parent.color == RED {
		if x.parent == x.grandparent().left {
			y = x.grandparent().right
			if y != nil && y.color == RED {
				x.parent.color = BLACK
				y.color = BLACK
				x.grandparent().color = RED
				x = x.grandparent()
			} else {
				if x == x.parent.right {
					x = x.parent
					t.leftRotate(x)
				}
				x.parent.color = BLACK
				x.grandparent().color = RED
				t.rightRotate(x.grandparent())
			}
		} else {
			y = x.grandparent().left
			if y != nil && y.color == RED {
				x.parent.color = BLACK
				y.color = BLACK
				x.grandparent().color = RED
				x = x.grandparent()
			} else {
				if x == x.parent.left {
					x = x.parent
					t.rightRotate(x)
				}
				x.parent.color = BLACK
				x.grandparent().color = RED
				t.leftRotate(x.grandparent())
			}
		}
	}

	t.root.color = BLACK
}

func (t *RedAndBlackTree) leftRotate(x *rbtNode) {
	y := x.right
	x.right = y.left
	if y.left != nil {
		y.left.parent = x
	}
	y.parent = x.parent

	if x.parent == nil {
		t.root = y
	} else {
		if x == x.parent.left {
			x.parent.left = y
		} else {
			x.parent.right = y
		}
	}
	y.left = x
	x.parent = y
}

func (t *RedAndBlackTree) rightRotate(x *rbtNode) {
	y := x.left
	x.left = y.right
	if y.right != nil {
		y.right.parent = x
	}
	y.parent = x.parent

	if x.parent == nil {
		t.root = y
	} else {
		if x == x.parent.right {
			x.parent.right = y
		} else {
			x.parent.left = y
		}
	}
	y.right = x
	x.parent = y
}

func (t *RedAndBlackTree) replace(a, b *rbtNode) {
	if a.parent == nil {
		t.root = b
	} else if a == a.parent.left {
		a.parent.left = b
	} else {
		a.parent.right = b
	}
	if b != nil {
		b.parent = a.parent
	}
}

func (t *RedAndBlackTree) Search(key int64) *rbtNode {
	x := t.root

	if x == nil {
		return nil
	}

	for x != nil {
		switch {
		case key == x.key:
			return x
		case key < x.key:
			x = x.left
		case key > x.key:
			x = x.right
		}
	}

	return nil
}

func (t *RedAndBlackTree) Delete(key int64) {
	z := t.Search(key)
	if z == nil {
		return
	}

	t.delete(z)
}

func (t *RedAndBlackTree) delete(z *rbtNode) *rbtNode {
	var x, y *rbtNode
	y = z

	if z.left == nil {
		x = z.right
		t.replace(z, z.right)
	} else if z.right == nil {
		x = z.left
		t.replace(z, z.left)

	} else {
		y = z.successor()
		if y.left != nil {
			x = y.left
		} else {
			x = y.right
		}
		x.parent = y.parent

		if y.parent == nil {
			t.root = x
		} else {
			if y == y.parent.left {
				y.parent.left = x
			} else {
				y.parent.right = x
			}
		}
	}

	if y.color == BLACK {
		t.deleteRepairNode(x)
	}
	t.size--

	return y
}

func (t *RedAndBlackTree) deleteRepairNode(x *rbtNode) {
	if x == nil {
		return
	}
	var w *rbtNode
	for x != t.root && x.color == BLACK {
		if x == x.parent.left {
			w = x.sibling()
			if w.color == RED {
				w.color = BLACK
				x.parent.color = RED
				t.leftRotate(x.parent)
				w = x.parent.right
			}
			if w.left.color == BLACK && w.right.color == BLACK {
				w.color = RED
				x = x.parent
			} else {
				if w.right.color == BLACK {
					w.left.color = BLACK
					w.color = RED
					t.rightRotate(w)
					w = x.parent.right
				}
				w.color = x.parent.color
				x.parent.color = BLACK
				w.right.color = BLACK
				t.leftRotate(x.parent)
				x = t.root
			}
		} else {
			w = x.sibling()
			if w.color == RED {
				w.color = BLACK
				x.parent.color = RED
				t.rightRotate(x.parent)
				w = x.parent.left
			}
			if w.left.color == BLACK && w.right.color == BLACK {
				w.color = RED
				x = x.parent
			} else {
				if w.left.color == BLACK {
					w.right.color = BLACK
					w.color = RED
					t.leftRotate(w)
					w = x.parent.left
				}
				w.color = x.parent.color
				x.parent.color = BLACK
				w.left.color = BLACK
				t.rightRotate(x.parent)
				x = t.root
			}

		}
	}
	x.color = BLACK
}

func (t *RedAndBlackTree) Size() int {
	return t.size
}

func (t *RedAndBlackTree) Minimum() *rbtNode {
	if t.root != nil {
		return t.root.minimum()
	}

	return nil
}

func newNode(key int64, value string) *rbtNode {
	return &rbtNode{
		key:   key,
		value: value,
	}
}

func (n *rbtNode) GetKey() int64 {
	return n.key
}

func (n *rbtNode) GetValue() string {
	return n.value
}

func (n *rbtNode) father() *rbtNode {
	return n.parent
}

func (n *rbtNode) grandparent() *rbtNode {
	g := n.father()
	if g == nil {
		return nil
	}

	return g.parent
}

func (n *rbtNode) sibling() *rbtNode {
	p := n.father()
	if p == nil {
		return nil
	}
	if n == p.left {
		return p.right
	}

	return p.left
}

func (n *rbtNode) successor() *rbtNode {
	if n.right != nil {
		return n.right.minimum()
	}

	y := n.parent
	for y != nil && n == y.right {
		n = y
		y = y.parent
	}

	return y
}

func (n *rbtNode) minimum() *rbtNode {
	for n.left != nil {
		n = n.left
	}

	return n
}

func (n *rbtNode) maximum() *rbtNode {
	for n.right != nil {
		n = n.right
	}
	return n
}

func (n *rbtNode) preorder() {
	fmt.Printf("(%v %v)", n.key, n.value)
	if n.parent == nil {
		fmt.Printf("nil")
	} else {
		fmt.Printf("whose parent is %v", n.parent.key)
	}
	if n.color == RED {
		fmt.Println(" and color RED")
	} else {
		fmt.Println(" and color BLACK")
	}
	if n.left != nil {
		fmt.Printf("%v's left child is ", n.key)
		n.left.preorder()
	}
	if n.right != nil {
		fmt.Printf("%v's right child is ", n.key)
		n.right.preorder()
	}
}
