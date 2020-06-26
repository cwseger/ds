package queue

// Queue is the interface that defines how to interact with the queue
type Queue interface {
	Enqueue(interface{})
	Dequeue() interface{}
	Peek() interface{}
	IsEmpty() bool
}

// queueElement is the element that is used to store data in the queue
type queueElement struct {
	value interface{}
	next  *queueElement
}

// DefaultQueue is the instance of the interface that can be used
type DefaultQueue struct {
	head *queueElement
	tail *queueElement
}

// New returns a new DefaultQueue
func New() *DefaultQueue {
	return &DefaultQueue{}
}

// Enqueue inserts the provided element into the back of the queue
func (q *DefaultQueue) Enqueue(i interface{}) {
	qe := &queueElement{
		value: i,
	}
	if q.tail != nil {
		q.tail.next = qe
	} else {
		q.head = qe
	}
	q.tail = qe
}

// Dequeue removes and returns the first added element from the front of the queue
func (q *DefaultQueue) Dequeue() interface{} {
	if q.head == nil {
		return nil
	}
	qe := q.head
	q.head = q.head.next
	if q.head == nil {
		q.tail = nil
	}
	return qe.value
}

// Peek returns the first added element from the front of the queue without removing it
func (q *DefaultQueue) Peek() interface{} {
	if q.head == nil {
		return nil
	}
	return q.head.value
}

// IsEmpty returns whether or not the queue is empty
func (q *DefaultQueue) IsEmpty() bool {
	return q.head == nil
}
