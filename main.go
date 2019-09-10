package main

import "sync"

//
// Goroutines #3
//

//
// Implement fixed size blocking queue, the queue must be thread-safe
//
// blocking queue is a queue that works same as normal queue but blocks enqueue/dequeue until the operation is available
//
// - to create a queue use function NewQueue(capacity int) *Queue, the capacity represents maximum elements of the queue
// - the queue should work with multiple goroutines, for example one routine enqueuing, two other routines are dequeueing
//

type Queue interface {

	// enqueue element into queue, if the size of the queue is >= capacity, the function blocks until the size is < capacity
	Enqueue(interface{})

	// dequeue element from queue, if the size is 0, the function blocks until the size > 0
	Dequeue() interface{}

	// returns current size of the queue
	Size() int

	// returns capacity of the queue
	Capacity() int

	// clear the queue, size will be 0
	Clear()
}

type MyQueue struct {
	size int
	cap  int
	ch   chan interface{}
}

func (q *MyQueue) Enqueue(obj interface{}) {
	q.ch <- obj
}

func (q *MyQueue) Dequeue() interface{} {
	return <-q.ch
}

func (q *MyQueue) Size() int {
	return len(q.ch)
}

func (q *MyQueue) Capacity() int {
	return q.cap
}

func (q *MyQueue) Clear() {
	close(q.ch)
	q.ch = make(chan interface{}, q.cap)
}

func NewQueue(cap int) Queue {
	return &MyQueue{
		size: 0,
		cap:  cap,
		ch:   make(chan interface{}, cap),
	}
}

func main() {
	q := NewQueue(10)

	iteration := 10000000
	wait := sync.WaitGroup{}
	wait.Add(3)

	go func() {
		defer wait.Done()
		for i := 0; i < iteration/2; i++ {
			q.Enqueue(i)
		}
	}()

	go func() {
		defer wait.Done()
		for i := 0; i < iteration/2; i++ {
			q.Enqueue(i)
		}
	}()

	go func() {
		defer wait.Done()
		for i := 0; i < iteration; i++ {
			q.Dequeue()
		}
	}()

	wait.Wait()
}
