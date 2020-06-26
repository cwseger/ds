package stack

// Stack is the interface that defines how to interact with the stack
type Stack interface {
	Push(interface{})
	Pop() interface{}
	Peek() interface{}
	IsEmpty() bool
}

// stackElement is the element that is used to store data on the stack
type stackElement struct {
	value interface{}
	next  *stackElement
}

// DefaultStack is the instance of the interface that can be used
type DefaultStack struct {
	top *stackElement
}

// New returns a new DefaultStack
func New() *DefaultStack {
	return &DefaultStack{}
}

// Push takes in an element and places it onto of the stack
func (s *DefaultStack) Push(i interface{}) {
	se := &stackElement{
		value: i,
		next:  s.top,
	}
	s.top = se
}

// Pop removes and returns the top element on the stack
func (s *DefaultStack) Pop() interface{} {
	if s.top == nil {
		return nil
	}
	se := s.top
	s.top = s.top.next
	return se.value
}

// Peek returns the top element on the stack without removing it
func (s *DefaultStack) Peek() interface{} {
	if s.top == nil {
		return nil
	}
	return s.top.value
}

// IsEmpty returns whether or not the stack contains any elements
func (s *DefaultStack) IsEmpty() bool {
	return s.top == nil
}
