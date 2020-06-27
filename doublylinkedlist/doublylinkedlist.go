package doublylinkedlist

type Node struct {
	Value    interface{}
	Next     *Node
	Previous *Node
}

type DoublyLinkedList struct {
	head *Node
}

func NewDoublyLinkedList() DoublyLinkedList {
	return DoublyLinkedList{}
}

func (dll *DoublyLinkedList) Add(value interface{}) {
	node := &Node{
		Value:    value,
		Next:     nil,
		Previous: dll.last(),
	}

	if dll.head == nil {
		dll.head = node
	} else {
		dll.last().Next = node
	}
}

func (dll *DoublyLinkedList) AddAll(values []interface{}) {
	last := dll.last()

	for _, value := range values {
		node := &Node{
			Value:    value,
			Next:     nil,
			Previous: last,
		}

		if last == nil {
			dll.head = node
		} else {
			last.Next = node
		}

		last = node
	}
}

func (dll *DoublyLinkedList) Get(index int) interface{} {
	current := dll.head
	for i := 0; i < index; i++ {
		current = current.Next
	}
	return current.Value
}

func (dll *DoublyLinkedList) Values() (values []interface{}) {
	current := dll.head
	for current != nil {
		values = append(values, current.Value)
		current = current.Next
	}
	return
}

func (dll *DoublyLinkedList) Clear() {
	dll.head = nil
}

func (dll *DoublyLinkedList) Contains(value interface{}) bool {
	return dll.IndexOf(value) != -1
}

// returns -1 if list does not contain the element
func (dll *DoublyLinkedList) IndexOf(value interface{}) (index int) {
	current := dll.head
	for current != nil {
		if current.Value == value {
			return
		}

		index++
		current = current.Next
	}

	return -1
}

func (dll *DoublyLinkedList) Size() (size int) {
	current := dll.head
	for current != nil {
		current = current.Next
		size++
	}
	return
}

func (dll *DoublyLinkedList) IsEmpty() bool {
	return dll.Size() == 0
}

func (dll *DoublyLinkedList) RemoveAt(index int) (success bool) {
	if index == 0 {
		if dll.head != nil {
			dll.head = dll.head.Next
			dll.head.Previous = nil
			success = true
		}
		return
	}

	current := dll.head
	for i := 0; i < index && current != nil; i++ {
		current = current.Next
	}

	if current != nil {
		if current.Previous != nil {
			current.Previous.Next = current.Next
		}

		if current.Next != nil {
			current.Next.Previous = current.Previous
		}

		success = true
	}

	return
}

// n is the maximum amount of removals to max, or -1 for unlimited
func (dll *DoublyLinkedList) Remove(value interface{}, n int) (removals int) {
	current := dll.head
	for current != nil && removals != n {
		if current.Value == value {
			if current.Next != nil {
				current.Next.Previous = current.Previous
			}

			if current.Previous == nil {
				dll.head = current.Next
			} else {
				current.Previous.Next = current.Next
			}

			removals++
		}

		current = current.Next
	}

	return
}

func (dll *DoublyLinkedList) Set(index int, value interface{}) {
	current := dll.head
	for i := 0; i < index; i++ {
		current = current.Next
	}

	current.Value = value
}

func (dll *DoublyLinkedList) last() *Node {
	if dll.head == nil {
		return nil
	}

	current := dll.head
	for current.Next != nil {
		current = current.Next
	}

	return current
}
