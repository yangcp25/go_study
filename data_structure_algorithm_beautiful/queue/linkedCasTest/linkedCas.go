package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"unsafe"
)

// Node represents a node in the queue.
type Node struct {
	value interface{}
	next  *Node
}

// CASQueue represents a concurrent queue using CAS operations.
type CASQueue struct {
	head *Node
	tail *Node
}

// NewCASQueue creates a new CASQueue.
func NewCASQueue() *CASQueue {
	// Create a dummy node to start the queue.
	dummy := &Node{}
	return &CASQueue{
		head: dummy,
		tail: dummy,
	}
}

// Enqueue adds a new element to the end of the queue.
func (q *CASQueue) Enqueue(value interface{}) {
	newNode := &Node{value: value}

	for {
		tail := q.tail
		next := tail.next

		// Ensure tail is still the current tail.
		if tail == q.tail {
			// If tail's next is nil, attempt to enqueue.
			if next == nil {
				if atomic.CompareAndSwapPointer((*unsafe.Pointer)(unsafe.Pointer(&tail.next)), unsafe.Pointer(next), unsafe.Pointer(newNode)) {
					// Enqueue succeeded, update the tail.
					atomic.CompareAndSwapPointer((*unsafe.Pointer)(unsafe.Pointer(&q.tail)), unsafe.Pointer(tail), unsafe.Pointer(newNode))
					return
				}
			} else {
				// Tail was not the actual tail, update the tail reference.
				atomic.CompareAndSwapPointer((*unsafe.Pointer)(unsafe.Pointer(&q.tail)), unsafe.Pointer(tail), unsafe.Pointer(next))
			}
		}
	}
}

// Dequeue removes and returns the front element from the queue.
// Returns nil if the queue is empty.
func (q *CASQueue) Dequeue() interface{} {
	for {
		head := q.head
		tail := q.tail
		first := head.next

		// Ensure head and tail are still the current head and tail.
		if head == q.head {
			// If head's next is nil, the queue is empty.
			if first == nil {
				return nil
			}

			// Attempt to dequeue.
			if head == tail {
				// The queue is not empty, but tail is not updated yet.
				// Try to update the tail to the actual tail.
				atomic.CompareAndSwapPointer((*unsafe.Pointer)(unsafe.Pointer(&q.tail)), unsafe.Pointer(tail), unsafe.Pointer(first))
			} else {
				// Try to dequeue the first element.
				if atomic.CompareAndSwapPointer((*unsafe.Pointer)(unsafe.Pointer(&q.head)), unsafe.Pointer(head), unsafe.Pointer(first)) {
					// Dequeue succeeded, return the value.
					value := first.value
					first.value = nil // Help with garbage collection
					return value
				}
			}
		}
	}
}

func main() {
	// Example usage of CASQueue
	queue := NewCASQueue()

	var wg sync.WaitGroup

	// Enqueue
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			queue.Enqueue(i)
		}
	}()

	// Dequeue
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			value := queue.Dequeue()
			fmt.Printf("%+v", value)
		}
	}()

	wg.Wait()
}
