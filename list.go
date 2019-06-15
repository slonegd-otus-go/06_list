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
		list.last = &item{Value: v, list: list}
		list.first = list.last
		return
	}
	item := &item{
		Value: v,
		prev:  list.last,
		list:  list,
	}
	list.last.next = item
	list.last = item
}

func (list *Type) PushFront(v interface{}) {
	list.len++

	if list.first == nil {
		list.first = &item{Value: v, list: list}
		list.last = list.first
		return
	}
	item := &item{
		Value: v,
		next:  list.first,
		list:  list,
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

func (list *Type) RemoveIf(predicate func(value interface{}) bool) {
	for item := list.first; item.next != nil; item = item.next {
		if predicate(item.Value) {
			item.Remove()
		}
	}
}

func (list *Type) String() string {
	var builder strings.Builder
	builder.WriteRune('[')
	item := list.first
	if item != nil {
		for ; item.next != nil; item = item.next {
			builder.WriteString(
				fmt.Sprintf("%v ", item.Value),
			)
		}
		builder.WriteString(
			fmt.Sprintf("%v", item.Value),
		)
	}
	builder.WriteRune(']')
	return builder.String()
}

type item struct {
	Value interface{}
	prev  *item
	next  *item
	list  *Type
}

func (item item) Next() *item {
	return item.next
}

func (item item) Prev() *item {
	return item.prev
}

func (item *item) Remove() {
	item.list.len--

	if item.next == nil && item.prev == nil {
		item.list.first = nil
		item.list.last = nil
		return
	}

	if item.list.first == item {
		item.list.first = item.next
		item.next.prev = nil
		return
	}

	if item.list.last == item {
		item.list.last = item.prev
		item.prev.next = nil
		return
	}

	item.prev.next = item.next
	item.next.prev = item.prev
}
