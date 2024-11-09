package queue

// QueueNode is a node in the queue.
type QueueNode[T any] struct {
	// elem is the element.
	elem T

	// next is the next node.
	next *QueueNode[T]
}

// Free implements common.Type.
func (node *QueueNode[T]) Free() {
	if node == nil {
		return
	}

	if node.next != nil {
		node.next.Free()
	}

	node.next = nil
}
