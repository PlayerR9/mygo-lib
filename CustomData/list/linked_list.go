package list

import (
	"sync"

	"github.com/PlayerR9/mygo-lib/CustomData/list/internal"
	"github.com/PlayerR9/mygo-lib/common"
	"github.com/PlayerR9/mygo-lib/mem"
)

// LinkedList is a simple implementation of a list that is backed by an linked list.
// This implementation is thread-safe.
//
// An empty linked list can be created using the `list := new(list.LinkedList[T])` constructor.
type LinkedList[T any] struct {
	// front is the front of the list.
	front *internal.ListNode[T]

	// back is the back of the list.
	back *mem.Ref[*internal.ListNode[T]]

	// front_mu is the mutex for the front of the list.
	front_mu sync.RWMutex

	// back_mu is the mutex for the back of the list.
	back_mu sync.RWMutex
}

// Enlist implements List.
//
// Never returns ErrFullList.
func (l *LinkedList[T]) Enlist(elem T) error {
	if l == nil {
		return common.ErrNilReceiver
	}

	l.back_mu.Lock()
	defer l.back_mu.Unlock()

	node := internal.NewListNode(elem)

	if l.back == nil {
		l.front_mu.Lock()
		defer l.front_mu.Unlock()

		l.front = node.Borrow()
		l.back = node
	} else {
		back := l.back.Borrow()

		_ = back.SetNext(node.Borrow())

		// WARNING: Lost ownership of back.
		l.back = node
	}

	return nil
}

// Prepend implements List.
//
// Never returns ErrFullList.
func (l *LinkedList[T]) Prepend(elem T) error {
	if l == nil {
		return common.ErrNilReceiver
	}

	l.front_mu.Lock()
	defer l.front_mu.Unlock()

	node := internal.NewListNode(elem)
	ptr := node.Borrow()

	_ = ptr.SetNext(l.front)

	if l.front == nil {
		l.back_mu.Lock()
		defer l.back_mu.Unlock()

		l.back = node
	}

	l.front = ptr

	return nil
}

// Delist implements List.
func (l *LinkedList[T]) Delist() (T, error) {
	if l == nil {
		return *new(T), common.ErrNilReceiver
	}

	l.front_mu.Lock()
	defer l.front_mu.Unlock()

	if l.front == nil {
		return *new(T), ErrEmptyList
	}

	front := l.front
	l.front = front.MustGetNext()
	_ = l.front.SetPrev(nil) // Clear the reference.
	_ = front.SetNext(nil)   // Clear the reference.

	if l.front == nil {
		l.back_mu.Lock()
		defer l.back_mu.Unlock()

		l.back = nil
	}

	elem := front.MustGetElem()
	return elem, nil
}

// Deback implements List.
func (l *LinkedList[T]) Deback() (T, error) {
	if l == nil {
		return *new(T), common.ErrNilReceiver
	}

	l.back_mu.Lock()
	defer l.back_mu.Unlock()

	if l.back == nil {
		return *new(T), ErrEmptyList
	}

	back := l.back
	l.back = back.MustGetPrev()
	_ = l.back.SetNext(nil) // Clear the reference.
	_ = back.SetPrev(nil)   // Clear the reference.

	if l.back == nil {
		l.front_mu.Lock()
		defer l.front_mu.Unlock()

		l.front = nil
	}

	elem := back.MustGetElem()
	return elem, nil
}

// Front implements List.
func (l *LinkedList[T]) Front() (T, error) {
	if l == nil {
		return *new(T), common.ErrNilReceiver
	}

	l.front_mu.RLock()
	defer l.front_mu.RUnlock()

	if l.front == nil {
		return *new(T), ErrEmptyList
	}

	elem := l.front.MustGetElem()
	return elem, nil
}

// Back implements List.
func (l *LinkedList[T]) Back() (T, error) {
	if l == nil {
		return *new(T), common.ErrNilReceiver
	}

	l.back_mu.RLock()
	defer l.back_mu.RUnlock()

	if l.back == nil {
		return *new(T), ErrEmptyList
	}

	elem := l.back.Borrow().MustGetElem()
	return elem, nil
}

// IsEmpty implements List.
func (l *LinkedList[T]) IsEmpty() bool {
	if l == nil {
		return true
	}

	l.front_mu.RLock()
	defer l.front_mu.RUnlock()

	return l.front == nil
}

// Size implements List.
func (l *LinkedList[T]) Size() uint {
	if l == nil {
		return 0
	}

	l.front_mu.RLock()
	defer l.front_mu.RUnlock()

	l.back_mu.Lock()
	defer l.back_mu.Unlock()

	var size uint

	for c := l.front; c != nil; c = c.MustGetNext() {
		size++
	}

	return size
}

// Reset implements common.Resetter.
func (l *LinkedList[T]) Reset() {
	if l == nil {
		return
	}

	l.front_mu.Lock()
	defer l.front_mu.Unlock()

	l.back_mu.Lock()
	defer l.back_mu.Unlock()

	if l.back != nil {
		err := mem.Free("l.back", l.back)
		if err != nil {
			panic(err)
		}

		l.back = nil
	}

	l.front = nil
}

// free is a private method that frees the list.
func (l *LinkedList[T]) free() error {
	if l == nil {
		return mem.ErrNilReceiver
	}

	l.front_mu.Lock()
	defer l.front_mu.Unlock()

	l.back_mu.Lock()
	defer l.back_mu.Unlock()

	if l.back != nil {
		err := mem.Free("l.back", l.back)
		if err != nil {
			return err
		}

		l.back = nil
	}

	l.front = nil

	return nil
}

// NewLinkedList creates a new LinkedList with the given elements.
//
// The given elements are added to the list in the order they are passed.
// The first element passed will be the front of the list, and the last
// element passed will be the back of the list.
//
// Parameters:
//   - elems: A variable number of elements to be added to the list.
//
// Returns:
//   - *LinkedList[T]: A pointer to a new LinkedList containing the given
//     elements.
func NewLinkedList[T any](elems ...T) *LinkedList[T] {
	list := new(LinkedList[T])

	if len(elems) == 0 {
		return list
	}

	slice := link_elements(elems)

	list.front = slice[0]
	list.back = slice[len(slice)-1]

	clear(slice)

	return list
}

// link_elements creates a slice of *ListNode from the given elements,
// linking each node to the next node in the slice. The last node in the
// slice will have no previous node.
//
// Parameters:
//   - elems: A slice of elements to be converted into list nodes.
//
// Returns:
//   - []*internal.ListNode[T]: A slice of linked list nodes, where each node
//     points to the next node as its previous node.
func link_elements[T any](elems []T) []*internal.ListNode[T] {
	slice := make([]*internal.ListNode[T], 0, len(elems))

	for _, elem := range elems {
		node := internal.NewListNode(elem)
		slice = append(slice, node)
	}

	for i := range slice[:len(slice)-1] {
		_ = slice[i].SetNext(slice[i+1])
	}

	return slice
}

// EnlistMany adds multiple elements to the list in the order they are passed.
//
// Parameters:
//   - elems: A slice of elements to be added to the list.
//
// Returns:
//   - uint: The number of elements successfully enlistd onto the list.
//   - error: An error of type common.ErrNilReceiver if the receiver is nil.
func (l *LinkedList[T]) EnlistMany(elems []T) (uint, error) {
	lenElems := uint(len(elems))
	if lenElems == 0 {
		return 0, nil
	} else if l == nil {
		return 0, common.ErrNilReceiver
	}

	l.back_mu.Lock()
	defer l.back_mu.Unlock()

	slice := link_elements(elems)

	if l.back == nil {
		l.front_mu.Lock()
		defer l.front_mu.Unlock()

		l.front = slice[0]
	} else {
		_ = l.back.SetNext(slice[0])
	}

	l.back = slice[len(slice)-1]

	clear(slice)

	return lenElems, nil
}

// PrependMany adds multiple elements at the start of the list in the order they are passed.
//
// Parameters:
//   - elems: A slice of elements to be added to the list.
//
// Returns:
//   - uint: The number of elements successfully prepended onto the list.
//   - error: An error of type common.ErrNilReceiver if the receiver is nil.
func (l *LinkedList[T]) PrependMany(elems []T) (uint, error) {
	lenElems := uint(len(elems))
	if lenElems == 0 {
		return 0, nil
	} else if l == nil {
		return 0, common.ErrNilReceiver
	}

	l.front_mu.Lock()
	defer l.front_mu.Unlock()

	slice := link_elements(elems)

	if l.front == nil {
		l.back_mu.Lock()
		defer l.back_mu.Unlock()

		l.back = slice[len(slice)-1]
	} else {
		_ = l.front.SetPrev(slice[len(slice)-1])

	}

	l.front = slice[0]
	clear(slice)

	return lenElems, nil
}
