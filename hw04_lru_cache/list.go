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
	items *ListItem
}

func (l *list) Len() int {
	var i int
	firstItem := l.Front()
	for firstItem != nil {
		firstItem = firstItem.Next
		i++
	}
	return i
}

func (l *list) Front() *ListItem {
	if l.items == nil {
		return nil
	}
	for l.items.Prev != nil {
		l.items = l.items.Prev
	}
	return l.items
}

func (l *list) Back() *ListItem {
	if l.items == nil {
		return nil
	}
	for l.items.Next != nil {
		l.items = l.items.Next
	}
	return l.items
}

func (l *list) PushFront(v interface{}) *ListItem {
	item := new(ListItem)
	item.Value = v
	if firstItem := l.Front(); firstItem == nil {
		l.items = item
	} else {
		item.Next = firstItem
		firstItem.Prev = item
		item.Prev = nil
	}
	return item
}

func (l *list) PushBack(v interface{}) *ListItem {
	item := new(ListItem)
	item.Value = v
	if lastItem := l.Back(); lastItem == nil {
		l.items = item
	} else {
		item.Prev = lastItem
		lastItem.Next = item
		item.Next = nil
	}
	return item
}

func (l *list) Remove(i *ListItem) {
	if i.Prev != nil {
		i.Prev.Next = i.Next
		l.items = i.Prev
	}
	if i.Next != nil {
		i.Next.Prev = i.Prev
		l.items = i.Next
	}
	if i.Next == nil && i.Prev == nil {
		l.items = nil
	}
}

func (l *list) MoveToFront(i *ListItem) {
	tmp := i.Value
	l.Remove(i)
	l.PushFront(tmp)
}
