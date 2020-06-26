package stack_test

import (
	"testing"

	"github.com/cwseger/ds/stack"
)

func TestNew(t *testing.T) {
	s := stack.New()
	if s == nil {
		t.Errorf("New() did not initialize stack. got: %v, want: %T", s, &stack.DefaultStack{})
	}
}

func TestPushAndPeek(t *testing.T) {
	s := stack.New()
	peekOut := s.Peek()
	if peekOut != nil {
		t.Errorf("Peek() did not return nil when stack is empty. got: %v, want: %v", peekOut, nil)
	}
	s.Push(10)
	peekOut = s.Peek()
	if peekOut != 10 {
		t.Errorf("Push() did not put element at the top of the stack. got: %v, want: %v", peekOut, 10)
	}
	s.Push(20)
	peekOut = s.Peek()
	if peekOut != 20 {
		t.Errorf("Push() did not put element at the top of the stack. got: %v, want: %v", peekOut, 20)
	}
}

func TestPop(t *testing.T) {
	s := stack.New()
	s.Push(10)
	s.Push(20)
	popOut := s.Pop()
	if popOut != 20 {
		t.Errorf("Pop() did not return top element of the stack. got: %v, want: %v", popOut, 20)
	}
	popOut = s.Pop()
	if popOut != 10 {
		t.Errorf("Pop() did not return top element of the stack. got: %v, want: %v", popOut, 10)
	}
	popOut = s.Pop()
	if popOut != nil {
		t.Errorf("Pop() did not return nil when no elements in stack. got: %v, want: %v", popOut, nil)
	}
}

func TestIsEmpty(t *testing.T) {
	s := stack.New()
	if !s.IsEmpty() {
		t.Errorf("IsEmpty() did not return true when no elements in stack. got: %v, want: %v", s.IsEmpty(), true)
	}
	s.Push(10)
	if s.IsEmpty() {
		t.Errorf("IsEmpty() did not return false when elements are in stack. got: %v, want: %v", s.IsEmpty(), false)
	}
}
