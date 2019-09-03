package main

import (
	"sync"
	"testing"
)

// test 1 producer - 1 consumer
func TestQueue(t *testing.T) {
	q := NewQueue(10)

	iteration := 10000000
	wait := sync.WaitGroup{}
	wait.Add(2)

	go func() {
		defer wait.Done()
		for i:=0; i<iteration; i++ {
			q.Enqueue(i)
		}
	}()

	go func() {
		defer wait.Done()
		for i:=0; i<iteration; i++ {
			q.Dequeue()
		}
	}()

	wait.Wait()
}

// test 2 producers - 1 consumer
func TestQueue2(t *testing.T) {
	q := NewQueue(10)

	iteration := 10000000
	wait := sync.WaitGroup{}
	wait.Add(3)

	go func() {
		defer wait.Done()
		for i:=0; i<iteration; i++ {
			q.Enqueue(i)
		}
	}()

	go func() {
		defer wait.Done()
		for i:=0; i<iteration/2; i++ {
			q.Dequeue()
		}
	}()

	go func() {
		defer wait.Done()
		for i:=0; i<iteration/2; i++ {
			q.Dequeue()
		}
	}()

	wait.Wait()
}

// test 1 producer - 2 consumers
func TestQueue3(t *testing.T) {
	q := NewQueue(10)

	iteration := 10000000
	wait := sync.WaitGroup{}
	wait.Add(3)

	go func() {
		defer wait.Done()
		for i:=0; i<iteration/2; i++ {
			q.Enqueue(i)
		}
	}()

	go func() {
		defer wait.Done()
		for i:=0; i<iteration/2; i++ {
			q.Enqueue(i)
		}
	}()

	go func() {
		defer wait.Done()
		for i:=0; i<iteration; i++ {
			q.Dequeue()
		}
	}()

	wait.Wait()
}
