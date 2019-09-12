package main

import (
	"fmt"
)

//
// Implement string HashSet
//
// - string HashSet is a Set that accepts string values
// - don't use map[] for implementation
// - HashSet is usually implemented as an array of buckets where each bucket contains a single linked list
// - https://en.wikipedia.org/wiki/Hash_table
//
// - initialize HashSet with bucket size = 1
// - implement all methods from the interface
// - all tests should pass
//
// - here's an example flow of adding an element:
// - 1) set.Add(<item>)
// - 2) calculate hash of the <item>
// - 3) use hash to get index of the bucket
// - 4) if the bucket is empty add your element as a first node of single linked list
// - 5) if the bucket is not empty (collision) append the element to the end of linked list
//

type HashSet interface {

	// Adds the specified element to a set, returns true if the element is added to the HashSet; false if the element is already present.
	Add(item string) bool

	// Removes the specified element from a HashSet, returns true if the element is successfully found and removed; otherwise false.
	Remove(item string) bool

	// Determines whether a HashSet object contains the specified element.
	Contains(item string) bool

	// Sets custom hash function.
	SetHashFunction(func(string) int)

	// Gets the number of elements that are contained in a HashSet.
	Count() int

	// Gets the number of buckets in a HashSet.
	BucketCount() int

	// Removes all elements from a HashSet.
	Clear()

	// Internal method that will resize and re-hash all elements if the load factor > 1 (load factor = number of elements / number of buckets)
	// resize operation will double the bucket size and rehash all elements
	// this should be called automatically when adding a new element
	resize()
}

type Bucket struct {
	node *Node
}

func (b Bucket) GetAllNode() []*Node {
	data := []*Node{}

	if b.node != nil {
		data = append(data, b.node)

		node := b.node.next
		for {
			if node == nil {
				break
			}

			data = append(data, node)
			node = node.next
		}
	}

	return data
}

func (b Bucket) GetDataFromAllNode() []string {
	data := []string{}

	if b.node != nil {
		data = append(data, b.node.data)

		node := b.node.next
		for {
			if node == nil {
				break
			}

			data = append(data, node.data)
			node = node.next
		}
	}

	return data
}

func (b Bucket) NodeCount() int {
	count := 0
	node := b.node
	for {
		if node == nil {
			break
		}
		count++
		node = node.next
	}

	return count
}

type Node struct {
	data string
	prev *Node
	next *Node
}

type MyHashSet struct {
	buckets  []Bucket
	hashFunc func(string) int
}

func (hs *MyHashSet) Add(item string) bool {
	bucket := hs.getBucket(item)

	if bucket.node != nil {

		nodes := bucket.GetAllNode()
		var lastNode *Node
		for _, v := range nodes {
			if v.data == item {
				//duplicate
				return false
			} else if v.next == nil {
				lastNode = v
			}
		}

		newNode := Node{
			data: item,
			prev: lastNode,
			next: nil,
		}

		lastNode.next = &newNode
	} else {
		bucket.node = &Node{
			data: item,
			prev: nil,
			next: nil,
		}
	}

	if hs.Count() > hs.BucketCount() {
		hs.resize()
	}

	return true
}

func (hs *MyHashSet) Remove(item string) bool {
	bucket := hs.getBucket(item)

	for _, v := range bucket.GetAllNode() {
		if v.data == item {
			if v.prev != nil {
				v.prev.next = v.next
			} else {
				bucket.node = nil
			}
			return true
		}
	}

	return false
}

func (hs *MyHashSet) Contains(item string) bool {
	bucket := hs.getBucket(item)

	for _, v := range bucket.GetAllNode() {
		if v.data == item {
			return true
		}
	}

	return false
}

func (hs *MyHashSet) SetHashFunction(function func(string) int) {
	hs.hashFunc = function
}

func (hs *MyHashSet) Count() int {
	count := 0
	for _, v := range hs.buckets {
		count += v.NodeCount()
	}
	return count
}

func (hs *MyHashSet) BucketCount() int {
	return len(hs.buckets)
}

func (hs *MyHashSet) Clear() {
	hs.buckets = make([]Bucket, initialBucketSize)
}

func (hs *MyHashSet) resize() {
	allData := hs.getAllData()
	size := hs.BucketCount() * 2
	hs.buckets = make([]Bucket, size)

	fmt.Println("resize ", size)

	for _, s := range allData {
		hs.Add(s)
	}
}

func (hs *MyHashSet) getAllData() []string {
	allData := []string{}
	for _, b := range hs.buckets {
		data := b.GetDataFromAllNode()
		if len(data) > 0 {
			allData = append(allData, data...)
		}
	}

	return allData
}

func (hs *MyHashSet) getBucket(item string) *Bucket {
	hash := hs.hashFunc(item)
	index := hash % hs.BucketCount()
	bucket := &hs.buckets[index]

	return bucket
}

const initialBucketSize = 1

func NewHashSet() HashSet {
	return &MyHashSet{
		buckets: make([]Bucket, initialBucketSize),
		hashFunc: func(string) int {
			return 0
		},
	}
}

func main() {
}
