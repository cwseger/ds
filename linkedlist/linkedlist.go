package linkedlist

// LinkedList is the interface that defines how to interact with the linked list.
type LinkedList interface {
	Append(v any)
	Prepend(v any)
	Traverse(func(v any))
}

// element is the element that is used to store data in the linked list.
type element struct {
	val  any
	prev *element
	next *element
}

// NewSingly returns a new SinglyLinkedList.
func NewSingly() *SinglyLinkedList {
	return &SinglyLinkedList{}
}

// SinglyLinkedList is an implementation of LinkedList that only maintains a head.
type SinglyLinkedList struct {
	head *element
}

// Append adds the provided value to the end of the linked list.
func (ll *SinglyLinkedList) Append(v any) {
	e := &element{val: v}
	if ll.head == nil {
		ll.head = e
	} else {
		cur := ll.head
		for cur.next != nil {
			cur = cur.next
		}
		cur.next = e
	}
}

// Prepend adds the provided value to the beginning of the linked list.
func (ll *SinglyLinkedList) Prepend(v any) {
	e := &element{val: v}
	if ll.head == nil {
		ll.head = e
	} else {
		e.next = ll.head
		ll.head = e
	}
}

// Traverse traverses the linked list, calling the provided func with each value.
func (ll *SinglyLinkedList) Traverse(f func(v any)) {
	cur := ll.head
	for cur != nil {
		f(cur.val)
		cur = cur.next
	}
}

// NewDoubly returns a new DoublyLinkedList.
func NewDoubly() *DoublyLinkedList {
	return &DoublyLinkedList{}
}

// DoublyLinkedList is an implementation of a LinkedList that maintains a head and tail.
type DoublyLinkedList struct {
	head *element
	tail *element
}

// Append adds the provided value to the end of the linked list.
func (ll *DoublyLinkedList) Append(v any) {
	e := &element{val: v}
	if ll.tail == nil {
		ll.head = e
		ll.tail = e
	} else {
		e.prev = ll.tail
		ll.tail.next = e
		ll.tail = e
	}
}

// Prepend adds the provided value to the beginning of the linked list.
func (ll *DoublyLinkedList) Prepend(v any) {
	e := &element{val: v}
	if ll.head == nil {
		ll.head = e
		ll.tail = e
	} else {
		ll.head.prev = e
		e.next = ll.head
		ll.head = e
	}
}

// TraverseForward traverses the linked list from head to tail, calling the provided func with each value.
func (ll *DoublyLinkedList) TraverseForward(f func(v any)) {
	cur := ll.head
	for cur != nil {
		f(cur.val)
		cur = cur.next
	}
}

// TraverseBackward traverses the linked list from tail to head, calling the provided func with each value.
func (ll *DoublyLinkedList) TraverseBackward(f func(v any)) {
	cur := ll.tail
	for cur != nil {
		f(cur.val)
		cur = cur.prev
	}
}
