package queue

// Queue is the interface that defines how to interact with the queue.
type Queue interface {
	Enqueue(v any)
	Dequeue() any
	Peek() any
	IsEmpty() bool
}

// element is the element that is used to store data in the queue.
type element struct {
	val  any
	next *element
}

// DefaultQueue is an implementation of Queue that uses a linked list as the backing data structure..
type DefaultQueue struct {
	head *element
	tail *element
}

// New returns a new DefaultQueue.
func New() *DefaultQueue {
	return &DefaultQueue{}
}

// Enqueue inserts the provided element into the back of the queue to perform queue operations.
func (q *DefaultQueue) Enqueue(v any) {
	e := &element{val: v}
	if q.tail != nil {
		q.tail.next = e
	} else {
		q.head = e
	}
	q.tail = e
}

// Dequeue removes and returns the element from the front of the queue.
func (q *DefaultQueue) Dequeue() any {
	if q.head == nil {
		return nil
	}
	qe := q.head
	q.head = q.head.next
	if q.head == nil {
		q.tail = nil
	}
	return qe.val
}

// Peek returns the element from the front of the queue without removing it.
func (q *DefaultQueue) Peek() any {
	if q.head == nil {
		return nil
	}
	return q.head.val
}

// IsEmpty returns whether or not the queue is empty.
func (q *DefaultQueue) IsEmpty() bool {
	return q.head == nil
}
