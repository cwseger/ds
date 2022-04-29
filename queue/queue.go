package queue

// Queue is the interface that defines how to interact with the queue.
type Queue interface {
	Enqueue(v any)
	Dequeue() any
	Peek() any
	IsEmpty() bool
}

// element is the element that is used to store data in the linked list queue.
type element struct {
	val  any
	next *element
}

// LinkedListQueue is an implementation of Queue that uses a linked list as the backing data structure.
type LinkedListQueue struct {
	head *element
	tail *element
}

// NewLinkedList returns a new LinkedListQueue.
func NewLinkedList() *LinkedListQueue {
	return &LinkedListQueue{}
}

// Enqueue inserts the provided element into the back of the queue.
func (q *LinkedListQueue) Enqueue(v any) {
	e := &element{val: v}
	if q.tail != nil {
		q.tail.next = e
	} else {
		q.head = e
	}
	q.tail = e
}

// Dequeue removes and returns the element from the front of the queue.
func (q *LinkedListQueue) Dequeue() any {
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
func (q *LinkedListQueue) Peek() any {
	if q.head == nil {
		return nil
	}
	return q.head.val
}

// IsEmpty returns whether or not the queue is empty.
func (q *LinkedListQueue) IsEmpty() bool {
	return q.head == nil
}

// SliceQueue is an implementation of Queue that uses a slice as the backing data structure..
type SliceQueue struct {
	s []any
}

// NewSlice returns a new SliceQueue.
func NewSlice() *SliceQueue {
	return &SliceQueue{}
}

// Enqueue inserts the provided element into the back of the queue.
func (q *SliceQueue) Enqueue(v any) {
	q.s = append(q.s, v)
}

// Dequeue removes and returns the element from the front of the queue.
func (q *SliceQueue) Dequeue() any {
	if len(q.s) == 0 {
		return nil
	}
	v := q.s[0]
	q.s = q.s[1:]
	return v
}

// Peek returns the element from the front of the queue without removing it.
func (q *SliceQueue) Peek() any {
	if len(q.s) == 0 {
		return nil
	}
	return q.s[0]
}

// IsEmpty returns whether or not the queue is empty.
func (q *SliceQueue) IsEmpty() bool {
	return len(q.s) == 0
}
