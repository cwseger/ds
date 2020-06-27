package linkedlist

type LinkedList interface{}

type DefaultLinkedList struct{}

func New() *DefaultLinkedList {
	return &DefaultLinkedList{}
}
