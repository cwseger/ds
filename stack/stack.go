package stack

// Stack is the interface that defines how to interact with the stack.
type Stack interface {
	Push(v any)
	Pop() any
	Peek() any
	IsEmpty() bool
}

// stackElement is the element that is used to store data on the stack.
type stackElement struct {
	val  any
	next *stackElement
}

// DefaultStack is an implementation of Stack that uses a linked list as the backing data structure.
type DefaultStack struct {
	top *stackElement
}

// New returns a new DefaultStack.
func New() *DefaultStack {
	return &DefaultStack{}
}

// Push takes in an element and places it onto of the stack.
func (s *DefaultStack) Push(v any) {
	se := &stackElement{
		val:  v,
		next: s.top,
	}
	s.top = se
}

// Pop removes and returns the top element on the stack.
func (s *DefaultStack) Pop() any {
	if s.top == nil {
		return nil
	}
	se := s.top
	s.top = s.top.next
	return se.val
}

// Peek returns the top element on the stack without removing it.
func (s *DefaultStack) Peek() any {
	if s.top == nil {
		return nil
	}
	return s.top.val
}

// IsEmpty returns whether or not the stack is empty.
func (s *DefaultStack) IsEmpty() bool {
	return s.top == nil
}
