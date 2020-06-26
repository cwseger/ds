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
	head *stackElement
}

// New returns a new DefaultStack
func New() *DefaultStack {
	return &DefaultStack{}
}

// Push takes in an element and places it onto of the stack
func (s *DefaultStack) Push(i interface{}) {
	se := &stackElement{
		value: i,
		next:  s.head,
	}
	s.head = se
}

// Pop will remove and return the top element on the stack
func (s *DefaultStack) Pop() interface{} {
	if s.head == nil {
		return nil
	}
	se := s.head
	s.head = s.head.next
	return se.value
}

// Peek will return the top element on the stack without removing it
func (s *DefaultStack) Peek() interface{} {
	if s.head == nil {
		return nil
	}
	return s.head.value
}

// IsEmpty returns whether or not the stack contains any elements
func (s *DefaultStack) IsEmpty() bool {
	return s.head == nil
}
