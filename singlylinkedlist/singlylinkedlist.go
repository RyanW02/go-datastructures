package singlylinkedlist

type Node struct {
	Value interface{}
	Next  *Node
}

type SinglyLinkedList struct {
	head *Node
}

func NewSinglyLinkedList() SinglyLinkedList {
	return SinglyLinkedList{}
}

func (sll *SinglyLinkedList) Add(value interface{}) {
	node := &Node{
		Value: value,
		Next:  nil,
	}

	if sll.head == nil {
		sll.head = node
	} else {
		sll.last().Next = node
	}
}

func (sll *SinglyLinkedList) AddAll(values []interface{}) {
	var next *Node
	for i := len(values) - 1; i >= 0; i-- {
		next = &Node{
			Value: values[i],
			Next:  next,
		}
	}

	if sll.head == nil {
		sll.head = next
	} else {
		sll.last().Next = next
	}
}

func (sll *SinglyLinkedList) Get(index int) interface{} {
	current := sll.head
	for i := 0; i < index; i++ {
		current = current.Next
	}
	return current.Value
}

func (sll *SinglyLinkedList) Values() (values []interface{}) {
	current := sll.head
	for current != nil {
		values = append(values, current.Value)
		current = current.Next
	}
	return
}

func (sll *SinglyLinkedList) Clear() {
	sll.head = nil
}

func (sll *SinglyLinkedList) Contains(value interface{}) bool {
	return sll.IndexOf(value) != -1
}

// returns -1 if list does not contain the element
func (sll *SinglyLinkedList) IndexOf(value interface{}) (index int) {
	current := sll.head
	for current != nil {
		if current.Value == value {
			return
		}

		index++
		current = current.Next
	}

	return -1
}

func (sll *SinglyLinkedList) Size() (size int) {
	current := sll.head
	for current != nil {
		current = current.Next
		size++
	}
	return
}

func (sll *SinglyLinkedList) IsEmpty() bool {
	return sll.Size() == 0
}

func (sll *SinglyLinkedList) RemoveAt(index int) (success bool) {
	if index == 0 {
		if sll.head != nil {
			sll.head = sll.head.Next
			success = true
		}
		return
	}

	previous := sll.head
	for i := 0; i < index-1 && previous != nil; i++ {
		previous = previous.Next
	}

	if previous != nil && previous.Next != nil {
		previous.Next = previous.Next.Next
		success = true
	}

	return
}

// n is the maximum amount of removals to max, or -1 for unlimited
func (sll *SinglyLinkedList) Remove(value interface{}, n int) (removals int) {
	if n == 0 || sll.head == nil {
		return
	}

	if sll.head != nil && sll.head.Value == value {
		n--
	}

	current := sll.head

	for current != nil && removals != n {
		for current.Next != nil && current.Next.Value == value {
			current.Next = current.Next.Next
			removals++
		}

		current = current.Next
	}

	if sll.head != nil && sll.head.Value == value {
		sll.head = sll.head.Next
		removals++
	}

	return
}

func (sll *SinglyLinkedList) Set(index int, value interface{}) {
	current := sll.head
	for i := 0; i < index; i++ {
		current = current.Next
	}

	current.Value = value

	return
}

func (sll *SinglyLinkedList) last() *Node {
	if sll.head == nil {
		return nil
	}

	current := sll.head
	for current.Next != nil {
		current = current.Next
	}

	return current
}
