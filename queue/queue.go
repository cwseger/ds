package queue

type Queue interface {
	Enqueue(interface{})
	Dequeue() interface{}
	Peek() interface{}
	IsEmpty() bool
}

type queueElement struct {
	value interface{}
	next  *queueElement
}

type DefaultQueue struct {
	head *queueElement
	tail *queueElement
}

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
	}
	q.tail = qe
}

// Dequeue removes and returns the first added element from the front of the queue
func (q *DefaultQueue) Dequeue() interface{} {
	if head == nil {
		return nil
	}
	qe := q.head
	q.head = q.head.next
	if q.head == nil {
		q.tail = nil
	}
	return qe.value
}

// Peek returns the first added element from the back of the queue without removing it
func (q *DefaultQueue) Peek() interface{} {
	if head == nil {
		return nil
	}
	return q.head.value
}

// IsEmpty returns whether or not the queue is empty
func (q *DefaultQueue) IsEmpty() bool {
	return head == nil
}
