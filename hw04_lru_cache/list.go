package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

func NewList() List {
	return new(list)
}

type list struct {
	front  *ListItem
	back   *ListItem
	length int
}

func (l list) Len() int {
	return l.length
}

func (l list) Front() *ListItem {
	return l.front
}

func (l list) Back() *ListItem {
	return l.back
}

func (l *list) PushFront(v interface{}) *ListItem {
	if l.front == nil {
		l.front = &ListItem{Value: v}
		l.back = l.front
	} else {
		l.front.Prev = &ListItem{Value: v, Next: l.front, Prev: nil}
		l.front = l.front.Prev
	}
	l.length++
	return l.front
}

func (l *list) PushBack(v interface{}) *ListItem {
	if l.front == nil {
		l.front = &ListItem{Value: v}
		l.back = l.front
	} else {
		l.back.Next = &ListItem{Value: v, Next: nil, Prev: l.back}
		l.back = l.back.Next
	}
	l.length++
	return l.back
}

func (l *list) Remove(i *ListItem) {
	if i == nil {
		return
	}
	if i.Next != nil {
		i.Next.Prev = i.Prev
	} else {
		l.back = i.Prev
	}
	if i.Prev != nil {
		i.Prev.Next = i.Next
	} else {
		l.front = i.Next
	}
	i.Next = nil
	i.Prev = nil
	l.length--
}

func (l *list) MoveToFront(i *ListItem) {
	tmp := i.Value
	l.Remove(i)
	l.PushFront(tmp)
}
