package stack

import "testing"

func TestPushPeek(t *testing.T) {
	stack := NewStack()

	wanted := 1
	stack.Push(1)

	got := stack.Peek()
	if got != wanted {
		t.Fatalf("failed pushpeek test. got: %v wanted: %v", got, wanted)
	}
}

func TestPushPop(t *testing.T) {
	stack := NewStack()

	wanted := 1
	stack.Push(1)

	got := stack.Pop()
	if got != wanted {
		t.Fatalf("failed pushpop test. got: %v wanted: %v", got, wanted)
	}

	if !stack.Empty() {
		t.Fatalf("failed pushpop test. got: %v wanted: %v", false, true)
	}
}

func TestSize(t *testing.T) {
	stack := NewStack()

	wanted := 100

	for i := 0; i < wanted; i++ {
		stack.Push(i)
	}

	got := stack.Size()
	if got != wanted {
		t.Fatalf("failed pushpop test. got: %v wanted: %v", got, wanted)
	}
}

func TestSearchValid(t *testing.T) {
	stack := NewStack()

	wanted := 50

	for i := 0; i < 100; i++ {
		stack.Push(i)
	}

	got := stack.Search(50)
	if got != wanted {
		t.Fatalf("failed searchvalid test. got: %v wanted: %v", got, wanted)
	}
}

func TestSearchInvalid(t *testing.T) {
	stack := NewStack()

	wanted := -1

	for i := 0; i < 100; i++ {
		stack.Push(i)
	}

	got := stack.Search(500)
	if got != wanted {
		t.Fatalf("failed searchinvalid test. got: %v wanted: %v", got, wanted)
	}
}
