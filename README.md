
# Datastructures  
Various datastructures implemented in Go

## [SinglyLinkedList](https://github.com/RyanW02/go-datastructures/tree/master/singlylinkedlist)
A traditional singly linked list, with each node holding a value and a pointer to the next node.

Recommended when you need to prioritise performance manipulating the list, rather than performance accessing the elements.

### Usage
```go
func main() {
	// instantiate a singly linked list
	sll := singlylinkedlist.NewSinglyLinkedList()

	// let's add some data
	sll.Add("hello")
	sll.AddAll([]interface{}{"world", "!"})
	sll.AddAll([]interface{}{"my_value", "my_value", "my_value",})

	// ... and remove some
	sll.RemoveAt(0)
	sll.Remove("my_value", -1) // n = max count to remove, when n < 0, unlimited

	// set the element at index 0 to "test"
	sll.Set(0, "test")

	// print all values in the list
	fmt.Println(sll.Values())

	// print whether the list contains a "!" element
	fmt.Println(sll.Contains("!"))

	// clear the list, print the size, and check whether it is empty
	sll.Clear()
	fmt.Println(sll.Size())
	fmt.Println(sll.IsEmpty())
}
```

### Testing
```bash
go test ./singlylinkedlist
```

## [DoublyLinkedList](https://github.com/RyanW02/go-datastructures/tree/master/doublylinkedlist)
A traditional doubly linked list, with each node holding a value, a pointer to the next node, and a pointer to the previous node.

Recommended when you need to prioritise performance manipulating the list, specifically performing better with removals than a singly linked list.

### Usage
```go
func main() {
	// instantiate a doubly linked list
	dll := doublylinkedlist.NewDoublyLinkedList()

	// let's add some data
	dll.Add("hello")
	dll.AddAll([]interface{}{"world", "!"})
	dll.AddAll([]interface{}{"my_value", "my_value", "my_value",})

	// ... and remove some
	dll.RemoveAt(0)
	dll.Remove("my_value", -1) // n = max count to remove, when n < 0, unlimited

	// set the element at index 0 to "test"
	dll.Set(0, "test")

	// print all values in the list
	fmt.Println(dll.Values())

	// print whether the list contains a "!" element
	fmt.Println(dll.Contains("!"))

	// clear the list, print the size, and check whether it is empty
	dll.Clear()
	fmt.Println(dll.Size())
	fmt.Println(dll.IsEmpty())
}
```

### Testing
```bash
go test ./doublylinkedlist
```

## [SliceList](https://github.com/RyanW02/go-datastructures/tree/master/slicelist)
A tribute to Java's ArrayList. Simply a wrapper around a slice, so you don't have to remember slice tricks off by heart.

Recommended when you need to prioritise performance in accessing data, rather than manipulating the list. However, you should probably just use a regular slice instead.

### Usage
```go
func main() {
	// instantiate a slice list
	sl := slicelist.NewSliceList()

	// let's add some data
	sl.Add("hello")
	sl.AddAll([]interface{}{"world", "!"})
	sl.AddAll([]interface{}{"my_value", "my_value", "my_value",})

	// ... and remove some
	sl.RemoveAt(0)
	sl.Remove("my_value", -1) // n = max count to remove, when n < 0, unlimited

	// set the element at index 0 to "test"
	sl.Set(0, "test")

	// print all values in the list
	fmt.Println(sl.Values())

	// print whether the list contains a "!" element
	fmt.Println(sl.Contains("!"))

	// clear the list, print the size, and check whether it is empty
	sl.Clear()
	fmt.Println(sl.Size())
	fmt.Println(sl.IsEmpty())
}
```

### Testing
```bash
go test ./slicelist
```

## [Stack](https://github.com/RyanW02/go-datastructures/tree/master/stack)
A traditional LIFO stack data structure.

### Usage
```go
func main() {
	// instantiate a stack
	stack := stack.NewStack()

	// let's add some data
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(1)
	stack.Push(3)

	// let's take a look at some data
	stack.Peek() // 1

	// check whether the stack is empty
	stack.Empty() // false
	
	// check the stack size
	stack.Size() // 5
	
	// pop values from the stack
	stack.Pop() // 1
	stack.Pop() // 2
	
	// search for the first instance of 3 in the stack
	stack.Search(3) // 0, since we popped 1 and 2
}
```

### Testing
```bash
go test ./stack
```