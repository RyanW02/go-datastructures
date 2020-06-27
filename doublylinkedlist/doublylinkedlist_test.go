package doublylinkedlist

import (
	"reflect"
	"testing"
)

func TestAddAll(t *testing.T) {
	dll := NewDoublyLinkedList()

	data := []interface{}{0,1,2,3}
	dll.AddAll(data)

	for i := 0; i < len(data); i++ {
		if value := dll.Get(i); value != data[i] {
			t.Fatalf("failed addall test. got: %v wanted: %v", value, 2)
		}
	}
}

func TestAddGet(t *testing.T) {
	dll := NewDoublyLinkedList()

	dll.AddAll([]interface{}{1,2})

	if value, ok := dll.Get(1).(int); !ok || value != 2 {
		t.Fatalf("failed add test. got: %v wanted: %v", value, 2)
	}
}

func TestValues(t *testing.T) {
	dll := NewDoublyLinkedList()

	data := []interface{}{0,1,2,2,3,2,4}
	dll.AddAll(data)

	got := dll.Values()
	if !reflect.DeepEqual(got, data) {
		t.Fatalf("failed values test. got: %v wanted: %v", got, data)
	}
}

func TestSize(t *testing.T) {
	dll := NewDoublyLinkedList()

	wanted := 100
	for i := 0; i < wanted; i++ {
		dll.Add(0)
	}

	got := dll.Size()
	if got != wanted {
		t.Fatalf("failed size test. got: %v wanted: %v", got, wanted)
	}
}

func TestEmpty(t *testing.T) {
	dll := NewDoublyLinkedList()

	if !dll.IsEmpty() {
		t.Fatalf("failed empty test. got: %v wanted: %v", false, true)
	}
}

func TestRemoveAtFirst(t *testing.T) {
	dll := NewDoublyLinkedList()
	dll.AddAll([]interface{}{0,1,2,3})

	// remove first element
	if !dll.RemoveAt(0) {
		t.Fatalf("failed remove first test. got: %v wanted: %v", false, true)
	}

	got := dll.Get(0)
	wanted := 1
	if got != wanted {
		t.Fatalf("failed removeat first test. got: %v wanted: %v", got, wanted)
	}
}

func TestRemoveAt(t *testing.T) {
	dll := NewDoublyLinkedList()
	dll.AddAll([]interface{}{0,1,2,3})

	// remove third element
	if !dll.RemoveAt(2) {
		t.Fatalf("failed remove first test. got: %v wanted: %v", false, true)
	}

	got := dll.Get(2)
	wanted := 3
	if got != wanted {
		t.Fatalf("failed removeat test. got: %v wanted: %v", got, wanted)
	}
}

func TestRemoveAtLast(t *testing.T) {
	dll := NewDoublyLinkedList()
	dll.AddAll([]interface{}{0,1,2,3})

	// remove last element
	if !dll.RemoveAt(dll.Size() - 1) {
		t.Fatalf("failed remove first test. got: %v wanted: %v", false, true)
	}

	got := dll.Get(dll.Size() - 1)
	wanted := 2
	if got != wanted {
		t.Fatalf("failed removeat last test. got: %v wanted: %v", got, wanted)
	}
}

func TestRemoveAtNop(t *testing.T) {
	dll := NewDoublyLinkedList()
	dll.AddAll([]interface{}{0,1,2,3})

	// remove last element
	if dll.RemoveAt(100) {
		t.Fatalf("failed removeat nop test. got: %v wanted: %v", true, false)
	}
}

func TestRemoveAll(t *testing.T) {
	dll := NewDoublyLinkedList()
	dll.AddAll([]interface{}{2,0,1,2,2,2,3,2,1,2,2,1,2})

	dll.Remove(2, -1)

	wanted := []interface{}{0,1,3,1,1}
	got := dll.Values()

	if !reflect.DeepEqual(got, wanted) {
		t.Fatalf("failed remove test. got: %v wanted: %v", got, wanted)
	}
}

func TestRemoveSome(t *testing.T) {
	dll := NewDoublyLinkedList()
	dll.AddAll([]interface{}{2,0,1,2,2,2,3,2,1,2,2,1,2})

	dll.Remove(2, 3)

	wanted := []interface{}{0,1,2,3,2,1,2,2,1,2}
	got := dll.Values()

	if !reflect.DeepEqual(got, wanted) {
		t.Fatalf("failed remove test. got: %v wanted: %v", got, wanted)
	}
}

func TestSet(t *testing.T) {
	dll := NewDoublyLinkedList()
	dll.AddAll([]interface{}{0,1,2,3})

	index := 2
	wanted := 100
	dll.Set(index, wanted)

	got := dll.Get(index)

	if got != wanted {
		t.Fatalf("failed remove first test. got: %v wanted: %v", got, wanted)
	}
}

func TestSetFirst(t *testing.T) {
	dll := NewDoublyLinkedList()
	dll.Add(0)
	dll.AddAll([]interface{}{0,1,2,3})

	index := 0
	wanted := 100
	dll.Set(index, wanted)

	got := dll.Get(index)

	if got != wanted {
		t.Fatalf("failed remove first test. got: %v wanted: %v", got, wanted)
	}
}

func TestSetFirstOneElement(t *testing.T) {
	dll := NewDoublyLinkedList()
	dll.Add(0)

	index := 0
	wanted := 100
	dll.Set(index, wanted)

	got := dll.Get(index)

	if got != wanted {
		t.Fatalf("failed remove first test. got: %v wanted: %v", got, wanted)
	}
}

func TestSetLast(t *testing.T) {
	dll := NewDoublyLinkedList()
	dll.AddAll([]interface{}{0,0,0,0})

	index := dll.Size() - 1
	wanted := 100
	dll.Set(index, wanted)

	got := dll.Get(index)

	if got != wanted {
		t.Fatalf("failed remove first test. got: %v wanted: %v", got, wanted)
	}
}

func TestIndexOf(t *testing.T) {
	dll := NewDoublyLinkedList()
	dll.AddAll([]interface{}{0,1,2,3})

	missingIndex := dll.IndexOf(4)
	if missingIndex != -1 {
		t.Fatalf("failed contains test. got: %v wanted: %v", missingIndex, -1)
	}

	validIndex := dll.IndexOf(1)
	if validIndex != 1 {
		t.Fatalf("failed contains test. got: %v wanted: %v", validIndex, 1)
	}
}

func TestContains(t *testing.T) {
	dll := NewDoublyLinkedList()
	dll.AddAll([]interface{}{0,1,2,3})

	if dll.Contains(4) {
		t.Fatalf("failed contains test. got: %v wanted: %v", true, false)
	}

	if !dll.Contains(3) {
		t.Fatalf("failed contains test. got: %v wanted: %v", false, true)
	}
}
