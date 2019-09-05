package main

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

	// Allow using custom hash function.
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

const initialBucketSize = 1

func NewHashSet() HashSet {
	return nil
}

func main() {
}
