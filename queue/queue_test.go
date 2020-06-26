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
