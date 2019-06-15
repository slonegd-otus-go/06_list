package list

import (
	"fmt"
	"strings"
)

type Type struct {
	first *item
	last  *item
	len   int
}

func (list Type) Len() int {
	return list.len
}

func (list *Type) PushBack(v interface{}) {
	list.len++

	if list.last == nil {
		list.last = &item{Value: v}
		list.first = list.last
		return
	}
	item := &item{
		Value: v,
		prev:  list.last,
	}
	list.last.next = item
	list.last = item
}

func (list *Type) PushFront(v interface{}) {
	list.len++

	if list.first == nil {
		list.first = &item{Value: v}
		list.last = list.first
		return
	}
	item := &item{
		Value: v,
		next:  list.first,
	}
	list.first.prev = item
	list.first = item
}

func (list Type) First() *item {
	return list.first
}

func (list Type) Last() *item {
	return list.last
}

func (list *Type) String() string {
	var builder strings.Builder
	builder.WriteRune('[')
	item := list.first
	for ; item.Next() != nil; item = item.Next() {
		builder.WriteString(
			fmt.Sprintf("%v ", item.Value),
		)
	}
	builder.WriteString(
		fmt.Sprintf("%v", item.Value),
	)
	builder.WriteRune(']')
	return builder.String()
}

type item struct {
	Value interface{}
	prev  *item
	next  *item
}

func (item item) Next() *item {
	return item.next
}

func (item item) Prev() *item {
	return item.prev
}
