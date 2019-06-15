package list

import (
	"fmt"
	"strings"
)

type List struct {
	first *item
	last  *item
	len   int
}

func (list List) Len() int {
	return list.len
}

func (list *List) PushBack(v interface{}) {
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

func (list *List) PushFront(v interface{}) {
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

func (list *List) PopBack() *item {
	return nil
}

func (list *List) PopFront() *item {
	return nil
}

func (list *List) Back() *item {
	return nil
}

func (list *List) Front() *item {
	return nil
}

func (list *List) Foreach(f func(*item)) {

}

func (list *List) FindIf(f func(*item) bool) *item {
	return nil
}

func (list *List) Remove(item *item) {

}

func (list *List) String() string {
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
