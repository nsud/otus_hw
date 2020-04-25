package hw04_lru_cache //nolint:golint,stylecheck

type List interface {
	Len() int                          // длина списка
	Front() *ListItem                  // первый Item
	Back() *ListItem                   // последний Item
	PushFront(v interface{}) *ListItem // добавить значение в начало
	PushBack(v interface{}) *ListItem  // добавить значение в конец
	Remove(i *ListItem)                // удалить элемент
	MoveToFront(i *ListItem)           // переместить элемент в начало
}

type ListItem struct {
	Value interface{} // значение
	Next  *ListItem   // следующий элемент
	Prev  *ListItem   // предыдущий элемент
}

type list struct {
	first *ListItem // первый элемент
	last  *ListItem // проследний элемент
	count int
}

func (l *list) Len() int {
	return l.count
}
func (l *list) Front() *ListItem {
	return l.first
}
func (l *list) Back() *ListItem {
	return l.last
}
func (l *list) PushFront(v interface{}) *ListItem { //в начало
	newItm := &ListItem{
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
	//fmt.Printf("PushFront: %v \t Next: %v \n", newItm.Value, newItm.Next)
	return newItm
}
func (l *list) PushBack(v interface{}) *ListItem { //в конец
	newItm := &ListItem{
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
	//fmt.Printf("PushBack: %v \t Next: %v \n", newItm.Value, newItm.Prev)

	return newItm
}
func (l *list) Remove(i *ListItem) {
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

func (l *list) MoveToFront(i *ListItem) {
	if l.Len() <= 1 || i == nil {
		return
	}
	l.Remove(i)
	l.PushFront(i.Value)
}

func NewList() List {
	return &list{}
}
