package queue_test

import (
	"testing"

	"github.com/cwseger/ds/queue"
)

func TestNewLinkedList(t *testing.T) {
	q := queue.NewLinkedList()
	if q == nil {
		t.Errorf("NewLinkedList() did not initialize queue. got: %v, want: %T", q, &queue.LinkedListQueue{})
	}
}

func TestLinkedListQueueEnqueueAndPeek(t *testing.T) {
	q := queue.NewLinkedList()
	peekOut := q.Peek()
	if peekOut != nil {
		t.Errorf("Peek() did not return nil when queue is empty. got: %v, want: %v", peekOut, nil)
	}
	q.Enqueue(10)
	peekOut = q.Peek()
	if peekOut != 10 {
		t.Errorf("Peek() did not return value of the element at the head. got: %v, want: %v", peekOut, 10)
	}
	q.Enqueue(20)
	peekOut = q.Peek()
	if peekOut != 10 {
		t.Errorf("Peek() did not return value of the element at the head. got: %v, want: %v", peekOut, 10)
	}
}

func TestLinkedListQueueDequeue(t *testing.T) {
	q := queue.NewLinkedList()
	q.Enqueue(10)
	q.Enqueue(20)
	dequeueOut := q.Dequeue()
	if dequeueOut != 10 {
		t.Errorf("Dequeue() did not return first element in the queue. got: %v, want: %v", dequeueOut, 10)
	}
	dequeueOut = q.Dequeue()
	if dequeueOut != 20 {
		t.Errorf("Dequeue() did not return first element in the queue. got: %v, want: %v", dequeueOut, 20)
	}
	dequeueOut = q.Dequeue()
	if dequeueOut != nil {
		t.Errorf("Dequeue() did not return nil when no elements in the queue. got: %v, want: %v", dequeueOut, nil)
	}
}

func TestLinkedListQueueIsEmpty(t *testing.T) {
	q := queue.NewLinkedList()
	if !q.IsEmpty() {
		t.Errorf("IsEmpty() did not return true when no elements in queue. got: %v, want: %v", q.IsEmpty(), true)
	}
	q.Enqueue(10)
	if q.IsEmpty() {
		t.Errorf("IsEmpty() did not return false when elements are in the queue. got: %v, want: %v", q.IsEmpty(), false)
	}
}

func TestNewSlice(t *testing.T) {
	q := queue.NewSlice()
	if q == nil {
		t.Errorf("NewSlice() did not initialize queue. got: %v, want: %T", q, &queue.SliceQueue{})
	}
}

func TestSliceQueueEnqueueAndPeek(t *testing.T) {
	q := queue.NewSlice()
	peekOut := q.Peek()
	if peekOut != nil {
		t.Errorf("Peek() did not return nil when queue is empty. got: %v, want: %v", peekOut, nil)
	}
	q.Enqueue(10)
	peekOut = q.Peek()
	if peekOut != 10 {
		t.Errorf("Peek() did not return value of the element at the head. got: %v, want: %v", peekOut, 10)
	}
	q.Enqueue(20)
	peekOut = q.Peek()
	if peekOut != 10 {
		t.Errorf("Peek() did not return value of the element at the head. got: %v, want: %v", peekOut, 10)
	}
}

func TestSliceQueueDequeue(t *testing.T) {
	q := queue.NewSlice()
	q.Enqueue(10)
	q.Enqueue(20)
	dequeueOut := q.Dequeue()
	if dequeueOut != 10 {
		t.Errorf("Dequeue() did not return first element in the queue. got: %v, want: %v", dequeueOut, 10)
	}
	dequeueOut = q.Dequeue()
	if dequeueOut != 20 {
		t.Errorf("Dequeue() did not return first element in the queue. got: %v, want: %v", dequeueOut, 20)
	}
	dequeueOut = q.Dequeue()
	if dequeueOut != nil {
		t.Errorf("Dequeue() did not return nil when no elements in the queue. got: %v, want: %v", dequeueOut, nil)
	}
}

func TestSliceQueueIsEmpty(t *testing.T) {
	q := queue.NewSlice()
	if !q.IsEmpty() {
		t.Errorf("IsEmpty() did not return true when no elements in queue. got: %v, want: %v", q.IsEmpty(), true)
	}
	q.Enqueue(10)
	if q.IsEmpty() {
		t.Errorf("IsEmpty() did not return false when elements are in the queue. got: %v, want: %v", q.IsEmpty(), false)
	}
}
