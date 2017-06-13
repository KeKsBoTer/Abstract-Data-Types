package twoThreeTree

import (
	"strconv"
	"fmt"
)

type tree struct {
	root node
}

/// Creates new EmptyTree
func NewTree() *tree {
	return &tree{}
}

func (t *tree) Insert(value int) {
	if t.root == nil {
		t.root = two{k: value}
	} else {
		t.root = t.root.insert(value)
	}
}

func (t *tree) String() string {
	if t.root != nil {
		return "{" + t.root.String() + "}"
	} else {
		return "{}"
	}
}

type node interface {
	ins(int) node
	insert(int) node
	grow() node
	restore() node
	String() string
}

type two struct {
	k           int
	left, right node
}

func newTwo(left node, k int, right node) *two {
	return &two{left: left, k: k, right: right}
}

type three struct {
	k1, k2           int
	left, mid, right node
}

type four struct {
	k1, k2, k3              int
	left, mid1, mid2, right node
}

func (t two) restore() node {
	if left, ok := t.left.(four); ok {
		return three{left: newTwo(left.left, left.k1, left.mid1), k1: left.k2, mid: newTwo(left.mid2, left.k3, left.right), k2: t.k, right: t.right}

	}
	if right, ok := t.left.(four); ok {
		return three{left: t.left, k1: t.k, mid: newTwo(right.left, right.k1, right.mid1), k2: right.k2, right: newTwo(right.mid2, right.k3, right.right)}
	}
	return t
}

func (t three) restore() node {
	if left, ok := t.left.(four); ok {
		return four{left: newTwo(left.left, left.k1, left.mid1), k1: left.k2, mid1: newTwo(left.mid2, left.k3, left.right), k2: t.k1, mid2: t.mid, k3: t.k2, right: t.right}
	}
	if mid, ok := t.left.(four); ok {
		return four{left: t.left, k1: t.k1, mid1: newTwo(mid.left, mid.k1, mid.mid1), k2: mid.k2, mid2: newTwo(mid.mid2, mid.k3, mid.right), k3: t.k2, right: t.right}
	}
	if right, ok := t.right.(four); ok {
		return four{left: t.left, k1: t.k1, mid1: t.mid, k2: t.k2, mid2: newTwo(right.left, right.k1, right.mid1), k3: right.k2, right: newTwo(right.mid2, right.k3, right.right)}
	}
	return t
}
func (t four) restore() node {
	return t
}

func (t two) ins(value int) node {
	switch {
	case t.k == value:
		return t
	case t.left == nil && t.right == nil:
		if t.k > value {
			return three{k1: value, k2: t.k}
		}
		if t.k < value {
			return three{k1: t.k, k2: value}
		}
	case t.left != nil && t.right != nil:
		if t.k > value {
			t.left = t.left.ins(value).restore()
		}
		if t.k < value {
			t.right = t.right.ins(value).restore()
		}
		return t
	}
	return nil
}

func (t three) ins(value int) node {
	switch {
	case t.k1 == value || t.k2 == value:
		return t
	case t.left == nil && t.mid == nil && t.right == nil:
		if value < t.k1 {
			return four{k1: value, k2: t.k1, k3: t.k2}
		} else if value < t.k2 {
			return four{k1: t.k1, k2: value, k3: t.k2}
		} else if value > t.k2 {
			return four{k1: t.k1, k2: t.k2, k3: value}
		}
		return t
	case t.left != nil && t.mid != nil && t.right != nil:
		if value < t.k1 {
			t.left = t.left.ins(value).restore()
		} else if value < t.k2 {
			t.mid = t.mid.ins(value)
		} else if value > t.k2 {
			a := t.right.ins(value)
			a = a.restore()
			t.right = a
		}
		return t
	}
	return nil
}

func (t four) ins(value int) node {
	println("error")
	return nil
}

func (t two) insert(value int) node {
	return t.ins(value).restore().grow()
}
func (t three) insert(value int) node {
	var a node= t.ins(value) //TODO undo
	fmt.Println(a)
	a = a.restore()
	fmt.Println(a)
	a = a.grow()
	fmt.Println(a)
	return a
}
func (t four) insert(value int) node {
	return t.ins(value).restore().grow()
}

func (t four) grow() node {
	return newTwo(newTwo(t.left, t.k1, t.mid1), t.k2, newTwo(t.mid2, t.k3, t.right))
}

func (t three) grow() node {
	return t
}

func (t two) grow() node {
	return t
}

func (t two) String() string {
	result := ""
	if t.left != nil {
		result += t.left.String() + ","
	}
	result += strconv.Itoa(t.k)
	if t.right != nil {
		result += "," + t.right.String()
	}
	return result
}

func (t three) String() string {
	result := ""
	if t.left != nil {
		result += t.left.String() + ","
	}
	result += strconv.Itoa(t.k1) + ","
	if t.mid != nil {
		result += t.mid.String() + ","
	}
	result += strconv.Itoa(t.k2)
	if t.right != nil {
		result += "," + t.right.String()
	}
	return result
}

func (t four) String() string {
	result := ""
	if t.left != nil {
		result += t.left.String() + ","
	}
	result += strconv.Itoa(t.k1) + ","
	if t.mid1 != nil {
		result += t.mid1.String() + ","
	}
	result += strconv.Itoa(t.k2) + ","
	if t.mid2 != nil {
		result += t.mid2.String() + ","
	}
	result += strconv.Itoa(t.k3)
	if t.right != nil {
		result += "," + t.right.String()
	}
	return result
}
