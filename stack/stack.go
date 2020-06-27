package stack

type Stack struct {
	slice []interface{}
}

func NewStack() Stack {
	return Stack{
		slice: make([]interface{}, 0),
	}
}

func (s *Stack) Empty() bool {
	return s.Size() == 0
}

func (s *Stack) Size() int {
	return len(s.slice)
}

func (s *Stack) Peek() interface{} {
	if s.Empty() {
		return nil
	}

	return s.slice[len(s.slice) - 1]
}

func (s *Stack) Pop() interface{} {
	value := s.Peek()
	s.slice = s.slice[:s.Size()-1]
	return value
}

func (s *Stack) Push(value interface{}) {
	s.slice = append(s.slice, value)
}

// returns -1 if not on the stack
func (s *Stack) Search(value interface{}) int {
	for i := s.Size() - 1; i >= 0; i-- {
		if s.slice[i] == value {
			return i
		}
	}

	return -1
}
