package queue_test

import (
	"testing"

	"github.com/cwseger/ds/queue"
)

func TestNew(t *testing.T) {
	q := queue.New()
	if q == nil {
		t.Errorf("New() did not initialize queue. got: %v, want: %T", q, &queue.DefaultQueue{})
	}
}

func TestEnqueueAndPeek(t *testing.T) {
	q := queue.New()
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

func TestDequeue(t *testing.T) {
	q := queue.New()
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

func TestIsEmpty(t *testing.T) {
	q := queue.New()
	if !q.IsEmpty() {
		t.Errorf("IsEmpty() did not return true when no elements in queue. got: %v, want: %v", q.IsEmpty(), true)
	}
	q.Enqueue(10)
	if q.IsEmpty() {
		t.Errorf("IsEmpty() did not return false when elements are in the queue. got: %v, want: %v", q.IsEmpty(), false)
	}
}
