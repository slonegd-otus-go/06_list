package list

import (
	"fmt"
	"strings"
)

type Node struct {
	Value interface{}
	prev  *Node
	next  *Node
}

func (node Node) Next() *Node {
	return node.next
}

type List struct {
	first *Node
	last  *Node
}

func (list *List) PushBack(v interface{}) {
	if list.last == nil {
		list.last = &Node{Value: v}
		list.first = list.last
		return
	}
	node := &Node{
		Value: v,
		prev:  list.last,
	}
	list.last.next = node
	list.last = node
}

func (list *List) PushFront(v interface{}) {
	if list.first == nil {
		list.first = &Node{Value: v}
		list.last = list.first
		return
	}
	node := &Node{
		Value: v,
		next:  list.first,
	}
	list.first.prev = node
	list.first = node
}

func (list *List) PopBack() *Node {
	return nil
}

func (list *List) PopFront() *Node {
	return nil
}

func (list *List) Back() *Node {
	return nil
}

func (list *List) Front() *Node {
	return nil
}

func (list *List) Foreach(f func(*Node)) {

}

func (list *List) FindIf(f func(*Node) bool) *Node {
	return nil
}

func (list *List) Remove(node *Node) {

}

func (list *List) String() string {
	var builder strings.Builder
	for node := list.first; node != nil; node = node.Next() {
		builder.WriteString(
			fmt.Sprintf("%v ", node.Value),
		)
	}
	return builder.String()
}
