package singlylinkedlist

import (
	"reflect"
	"testing"
)

func TestAddAll(t *testing.T) {
	sll := NewSinglyLinkedList()

	data := []interface{}{0,1,2,3}
	sll.AddAll(data)

	for i := 0; i < len(data); i++ {
		if value := sll.Get(i); value != data[i] {
			t.Fatalf("failed addall test. got: %v wanted: %v", value, 2)
		}
	}
}

func TestAddGet(t *testing.T) {
	sll := NewSinglyLinkedList()

	sll.AddAll([]interface{}{1,2})

	if value, ok := sll.Get(1).(int); !ok || value != 2 {
		t.Fatalf("failed add test. got: %v wanted: %v", value, 2)
	}
}

func TestValues(t *testing.T) {
	sll := NewSinglyLinkedList()

	data := []interface{}{0,1,2,2,3,2,4}
	sll.AddAll(data)

	got := sll.Values()
	if !reflect.DeepEqual(got, data) {
		t.Fatalf("failed values test. got: %v wanted: %v", got, data)
	}
}

func TestSize(t *testing.T) {
	sll := NewSinglyLinkedList()

	wanted := 100
	for i := 0; i < wanted; i++ {
		sll.Add(0)
	}

	got := sll.Size()
	if got != wanted {
		t.Fatalf("failed size test. got: %v wanted: %v", got, wanted)
	}
}

func TestEmpty(t *testing.T) {
	sll := NewSinglyLinkedList()

	if !sll.IsEmpty() {
		t.Fatalf("failed empty test. got: %v wanted: %v", false, true)
	}
}

func TestRemoveAtFirst(t *testing.T) {
	sll := NewSinglyLinkedList()
	sll.AddAll([]interface{}{0,1,2,3})

	// remove first element
	if !sll.RemoveAt(0) {
		t.Fatalf("failed remove first test. got: %v wanted: %v", false, true)
	}

	got := sll.Get(0)
	wanted := 1
	if got != wanted {
		t.Fatalf("failed removeat first test. got: %v wanted: %v", got, wanted)
	}
}

func TestRemoveAt(t *testing.T) {
	sll := NewSinglyLinkedList()
	sll.AddAll([]interface{}{0,1,2,3})

	// remove third element
	if !sll.RemoveAt(2) {
		t.Fatalf("failed remove first test. got: %v wanted: %v", false, true)
	}

	got := sll.Get(2)
	wanted := 3
	if got != wanted {
		t.Fatalf("failed removeat test. got: %v wanted: %v", got, wanted)
	}
}

func TestRemoveAtLast(t *testing.T) {
	sll := NewSinglyLinkedList()
	sll.AddAll([]interface{}{0,1,2,3})

	// remove last element
	if !sll.RemoveAt(sll.Size() - 1) {
		t.Fatalf("failed remove first test. got: %v wanted: %v", false, true)
	}

	got := sll.Get(sll.Size() - 1)
	wanted := 2
	if got != wanted {
		t.Fatalf("failed removeat last test. got: %v wanted: %v", got, wanted)
	}
}

func TestRemoveAtNop(t *testing.T) {
	sll := NewSinglyLinkedList()
	sll.AddAll([]interface{}{0,1,2,3})

	// remove last element
	if sll.RemoveAt(100) {
		t.Fatalf("failed removeat nop test. got: %v wanted: %v", true, false)
	}
}

func TestRemoveAll(t *testing.T) {
	sll := NewSinglyLinkedList()
	sll.AddAll([]interface{}{2,0,1,2,2,2,3,2,1,2,2,1,2})

	sll.Remove(2, -1)

	wanted := []interface{}{0,1,3,1,1}
	got := sll.Values()

	if !reflect.DeepEqual(got, wanted) {
		t.Fatalf("failed remove test. got: %v wanted: %v", got, wanted)
	}
}

func TestRemoveSome(t *testing.T) {
	sll := NewSinglyLinkedList()
	sll.AddAll([]interface{}{2,0,1,2,2,2,3,2,1,2,2,1,2})

	sll.Remove(2, 3)

	wanted := []interface{}{0,1,2,3,2,1,2,2,1,2}
	got := sll.Values()

	if !reflect.DeepEqual(got, wanted) {
		t.Fatalf("failed remove test. got: %v wanted: %v", got, wanted)
	}
}

func TestSet(t *testing.T) {
	sll := NewSinglyLinkedList()
	sll.AddAll([]interface{}{0,1,2,3})

	index := 2
	wanted := 100
	sll.Set(index, wanted)

	got := sll.Get(index)

	if got != wanted {
		t.Fatalf("failed remove first test. got: %v wanted: %v", got, wanted)
	}
}

func TestSetFirst(t *testing.T) {
	sll := NewSinglyLinkedList()
	sll.Add(0)
	sll.AddAll([]interface{}{0,1,2,3})

	index := 0
	wanted := 100
	sll.Set(index, wanted)

	got := sll.Get(index)

	if got != wanted {
		t.Fatalf("failed remove first test. got: %v wanted: %v", got, wanted)
	}
}

func TestSetFirstOneElement(t *testing.T) {
	sll := NewSinglyLinkedList()
	sll.Add(0)

	index := 0
	wanted := 100
	sll.Set(index, wanted)

	got := sll.Get(index)

	if got != wanted {
		t.Fatalf("failed remove first test. got: %v wanted: %v", got, wanted)
	}
}

func TestSetLast(t *testing.T) {
	sll := NewSinglyLinkedList()
	sll.AddAll([]interface{}{0,0,0,0})

	index := sll.Size() - 1
	wanted := 100
	sll.Set(index, wanted)

	got := sll.Get(index)

	if got != wanted {
		t.Fatalf("failed remove first test. got: %v wanted: %v", got, wanted)
	}
}

func TestIndexOf(t *testing.T) {
	sll := NewSinglyLinkedList()
	sll.AddAll([]interface{}{0,1,2,3})

	missingIndex := sll.IndexOf(4)
	if missingIndex != -1 {
		t.Fatalf("failed contains test. got: %v wanted: %v", missingIndex, -1)
	}

	validIndex := sll.IndexOf(1)
	if validIndex != 1 {
		t.Fatalf("failed contains test. got: %v wanted: %v", validIndex, 1)
	}
}

func TestContains(t *testing.T) {
	sll := NewSinglyLinkedList()
	sll.AddAll([]interface{}{0,1,2,3})

	if sll.Contains(4) {
		t.Fatalf("failed contains test. got: %v wanted: %v", true, false)
	}

	if !sll.Contains(3) {
		t.Fatalf("failed contains test. got: %v wanted: %v", false, true)
	}
}
