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

type list struct {
	length int
	front  *ListItem
	back   *ListItem
}

func NewList() List {
	return new(list)
}

func (l *list) Len() int {
	return l.length
}

func (l *list) Front() *ListItem {
	return l.front
}

func (l *list) Back() *ListItem {
	return l.back
}

func (l *list) PushFront(value interface{}) *ListItem {
	newListItem := &ListItem{Value: value, Next: l.front}

	if l.front != nil {
		l.front.Prev = newListItem
	}

	l.front = newListItem

	if l.back == nil {
		l.back = newListItem
	}

	l.length++
	return newListItem
}

func (l *list) PushBack(value interface{}) *ListItem {
	newListItem := &ListItem{Value: value, Prev: l.back}

	if l.back != nil {
		l.back.Next = newListItem
	}

	l.back = newListItem

	if l.front == nil {
		l.front = newListItem
	}

	l.length++
	return newListItem
}

func (l *list) Remove(li *ListItem) {
	if li.Prev != nil {
		li.Prev.Next = li.Next
	} else {
		l.front = li.Next
	}

	if li.Next != nil {
		li.Next.Prev = li.Prev
	} else {
		l.back = li.Prev
	}

	l.length--
}

func (l *list) MoveToFront(li *ListItem) {
	if l.front != li {
		l.Remove(li)
		l.PushFront(li.Value)
	}
}
