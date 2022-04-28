package linkedlist_test

import (
	"testing"

	"github.com/cwseger/ds/linkedlist"
)

func TestNewSingly(t *testing.T) {
	ll := linkedlist.NewSingly()
	if ll == nil {
		t.Errorf("NewSingly() did not initialize linked list. got: %v, want: %T", ll, &linkedlist.SinglyLinkedList{})
	}
}

func TestSinglyAppendPrependAndTraverse(t *testing.T) {
	ll := linkedlist.NewSingly()
	traverseCalledTimes := 0
	ll.Traverse(func(e any) {
		traverseCalledTimes++
	})
	if traverseCalledTimes != 0 {
		t.Errorf("Traverse() returned element when linked list is empty.")
	}
	ll.Append(10)
	ll.Append(20)
	ll.Prepend(30)
	ll.Prepend(40)
	ll.Traverse(func(e any) {
		traverseCalledTimes++
		switch traverseCalledTimes {
		case 1:
			if e.(int) != 40 {
				t.Errorf("Traverse() returned incorrect element. got: %v, want: %v", e, 40)
			}
		case 2:
			if e.(int) != 30 {
				t.Errorf("Traverse() returned incorrect element. got: %v, want: %v", e, 30)
			}
		case 3:
			if e.(int) != 10 {
				t.Errorf("Traverse() returned incorrect element. got: %v, want: %v", e, 10)
			}
		case 4:
			if e.(int) != 20 {
				t.Errorf("Traverse() returned incorrect element. got: %v, want: %v", e, 20)
			}
		}
	})
	if traverseCalledTimes != 4 {
		t.Errorf("Traverse() returned incorrect number of elements.")
	}
}

func TestNewDoubly(t *testing.T) {
	ll := linkedlist.NewDoubly()
	if ll == nil {
		t.Errorf("NewDoubly() did not initialize linked list. got: %v, want: %T", ll, &linkedlist.DoublyLinkedList{})
	}
}

func TestDoublyAppendAndTraverse(t *testing.T) {
	ll := linkedlist.NewDoubly()
	traverseForwardCalledTimes := 0
	traverseBackwardCalledTimes := 0
	ll.TraverseForward(func(e any) {
		traverseForwardCalledTimes++
	})
	if traverseForwardCalledTimes != 0 {
		t.Errorf("TraverseForward() returned element when linked list is empty.")
	}
	ll.TraverseBackward(func(e any) {
		traverseBackwardCalledTimes++
	})
	if traverseBackwardCalledTimes != 0 {
		t.Errorf("TraverseBackward() returned element when linked list is empty.")
	}
	ll.Append(10)
	ll.Append(20)
	ll.Prepend(30)
	ll.Prepend(40)
	ll.TraverseForward(func(e any) {
		traverseForwardCalledTimes++
		switch traverseForwardCalledTimes {
		case 1:
			if e.(int) != 40 {
				t.Errorf("TraverseForward() returned incorrect element. got: %v, want: %v", e, 40)
			}
		case 2:
			if e.(int) != 30 {
				t.Errorf("TraverseForward() returned incorrect element. got: %v, want: %v", e, 30)
			}
		case 3:
			if e.(int) != 10 {
				t.Errorf("TraverseForward() returned incorrect element. got: %v, want: %v", e, 10)
			}
		case 4:
			if e.(int) != 20 {
				t.Errorf("TraverseForward() returned incorrect element. got: %v, want: %v", e, 20)
			}
		}
	})
	if traverseForwardCalledTimes != 4 {
		t.Errorf("TraverseForward() returned incorrect number of elements.")
	}

	ll.TraverseBackward(func(e any) {
		traverseBackwardCalledTimes++
		switch traverseBackwardCalledTimes {
		case 1:
			if e.(int) != 20 {
				t.Errorf("TraverseBackward() returned incorrect element. got: %v, want: %v", e, 40)
			}
		case 2:
			if e.(int) != 10 {
				t.Errorf("TraverseBackward() returned incorrect element. got: %v, want: %v", e, 30)
			}
		case 3:
			if e.(int) != 30 {
				t.Errorf("TraverseBackward() returned incorrect element. got: %v, want: %v", e, 10)
			}
		case 4:
			if e.(int) != 40 {
				t.Errorf("TraverseBackward() returned incorrect element. got: %v, want: %v", e, 20)
			}
		}
	})
	if traverseBackwardCalledTimes != 4 {
		t.Errorf("TraverseBackward() returned incorrect number of elements.")
	}
}
