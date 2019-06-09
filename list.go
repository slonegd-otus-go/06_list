package list

type Node struct {
}

type List struct {
}

func (list *List) PushBack(v interface{}) {

}

func (list *List) PushFront(v interface{}) {

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
	return ""
}
