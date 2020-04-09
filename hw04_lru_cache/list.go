package hw04_lru_cache //nolint:golint,stylecheck

type List interface {
	Len() int                          // длина списка
	Front() *listItem                  // первый Item
	Back() *listItem                   // последний Item
	PushFront(v interface{}) *listItem // добавить значение в начало
	PushBack(v interface{}) *listItem  // добавить значение в конец
	Remove(i *listItem)                // удалить элемент
	MoveToFront(i *listItem)           // переместить элемент в начало
}

type listItem struct {
	Value interface{} // значение
	Next  *listItem   // следующий элемент
	Prev  *listItem   // предыдущий элемент
}

type list struct {
	first *listItem // первый элемент
	last  *listItem // проследний элемент
	count int
}

func (l *list) Len() int {
	return l.count
}
func (l *list) Front() *listItem {
	return l.first
}
func (l *list) Back() *listItem {
	return l.last
}
func (l *list) PushFront(v interface{}) *listItem { //в начало
	newItm := &listItem{
		Value: v,
		Next:  nil,
		Prev:  l.first,
	}
	if l.first == nil {
		l.last = newItm
	} else {
		l.first.Next = newItm
	}
	l.first = newItm
	l.count++
	return newItm
}
func (l *list) PushBack(v interface{}) *listItem { //в конец
	newItm := &listItem{
		Value: v,
		Next:  l.last,
		Prev:  nil,
	}
	if l.last == nil {
		l.first = newItm
	} else {
		l.last.Prev = newItm
	}
	l.last = newItm
	l.count++

	return newItm
}
func (l *list) Remove(i *listItem) {
	if i.Value == nil {
		return
	}
	if i.Next != nil {
		i.Next.Prev = i.Prev
	} else {
		l.first = i.Prev
	}
	if i.Prev != nil {
		i.Prev.Next = i.Next
	} else {
		l.last = i.Next
	}
	l.count--
}

func (l *list) MoveToFront(i *listItem) {
	if l.Len() <= 1 || i == nil {
		return
	}
	l.Remove(i)
	l.PushFront(i.Value)
}

func NewList() List {
	return &list{}
}
