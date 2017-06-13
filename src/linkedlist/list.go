package linkedlist

import "fmt"

type List struct {
	root *chain
}

func NewList() *List {
	return &List{}
}

func (l *List) Insert(key, value interface{}) {
	if l.root == nil {
		l.root = &chain{key:key,value: value}
	} else {
		l.root.insert(key, value)
	}
}

func (l *List) Delete(value interface{}) bool {
	if l.root == nil {
		return false
	} else {
		ok := false
		l.root, ok = l.root.delete(value)
		return ok
	}
}

func (l *List) Find(key interface{}) interface{} {
	if l.root == nil {
		return nil
	} else {
		return l.root.find(key)
	}
}

func (l *List) Length() int {
	if l.root == nil {
		return 0
	} else {
		return l.root.length()
	}
}

func (l *List) Contains(value interface{}) bool {
	if l.root == nil {
		return false
	} else {
		return l.root.contains(value)
	}
}

func (l *List) ForEach(f func(interface{}, interface{})) {
	if l.root != nil {
		l.root.iterate(f)
	}
}

func (l *List) String() string {
	if l.root == nil {
		return "[]"
	} else {
		return "[" + l.root.String() + "]"
	}
}

type chain struct {
	key, value interface{}
	next       *chain
}

func (c *chain) insert(key, value interface{}) {
	if c.next == nil {
		c.next = &chain{key: key, value: value}
	} else {
		c.next.insert(key, value)
	}
}

func (c *chain) delete(key interface{}) (*chain, bool) {
	if c.value == key {
		return c.next, true
	} else if c.next != nil {
		ok := false
		c.next, ok = c.next.delete(key)
		return c, ok
	} else {
		return c, false
	}
}

func (c *chain) find(key interface{}) interface{} {
	if c.key == key {
		return c.value
	} else if c.next != nil {
		return c.next.find(key)
	} else {
		return nil
	}
}

func (c *chain) length() int {
	if c.next == nil {
		return 1
	} else {
		return 1 + c.next.length()
	}
}

func (c *chain) contains(key interface{}) bool {
	if c.value == key {
		return true
	} else {
		return c.next != nil && c.next.contains(key)
	}
}

func (c *chain) String() string {
	result := fmt.Sprintf("%v", c.key) + "=" + fmt.Sprintf("%v", c.value)
	if c.next != nil {
		result += "," + c.next.String()
	}
	return result
}

func (c *chain) iterate(f func(interface{}, interface{})) {
	f(c.key, c.value)
	if c.next != nil {
		c.iterate(f)
	}
}
