package main

import (
	"gotest.tools/assert"
	"math/rand"
	"strconv"
	"testing"
)

func TestSet_Add(t *testing.T) {

	set := NewHashSet()

	// use crappy hash function to always collide, resize should be triggered when count == power^2
	set.SetHashFunction(func(s string) int {
		return 1
	})

	for i := 1; i<=128; i++ {
		item := strconv.Itoa(i)
		assert.Assert(t, !set.Contains(item), "Item %s can't be inside!", item)

		set.Add(item)
		assert.Assert(t, set.Contains(item), "Item %s must be inside!", item)
		assert.Assert(t, set.Count() == i, "Invalid set count: %d", set.Count())

		isPow := (set.Count() & (set.Count() - 1)) == 0

		// resize should be triggered when count == power^2
		if isPow {
			assert.Assert(t, set.BucketCount() == set.Count(), "Invalid bucket count: %d, expected: %d", set.BucketCount(), set.Count())
		}
	}
}

func TestSet_Dup(t *testing.T) {

	set := NewHashSet()
	set.SetHashFunction(func(s string) int {
		return rand.Int()
	})

	assert.Assert(t, set.Add("item"), "Item should be added!")
	assert.Assert(t, !set.Add("item"), "Item cannot be added!")
	assert.Assert(t, set.Count() == 1, "Item count must be 1!")
}

func TestSet_Remove(t *testing.T) {

	set := NewHashSet()
	set.SetHashFunction(func(s string) int {
		return rand.Int()
	})

	counter := 0

	for i := 1; i<=128; i++ {
		item := strconv.Itoa(i)
		assert.Assert(t, !set.Contains(item), "Item %s can't be inside!", item)

		set.Add(item)
		counter ++

		assert.Assert(t, set.Contains(item), "Item %s must be inside!", item)
		assert.Assert(t, set.Count() == counter, "Invalid set count: %d, expected: %d", set.Count(), counter)

		removed := set.Remove(item)
		counter --
		assert.Assert(t, removed, "Item %s must be removed!", item)
		assert.Assert(t, set.Count() == counter, "Invalid set count: %d, expected: %d", set.Count(), counter)

		assert.Assert(t, !set.Remove(item), "Item cannot be removed twice!")
		assert.Assert(t, set.Count() == counter, "Invalid set count: %d, expected: %d", set.Count(), counter)
	}

	assert.Assert(t, set.Count() == 0, "All elements must be 0")
}

func TestSet_Clear(t *testing.T) {

	set := NewHashSet()
	set.SetHashFunction(func(s string) int {
		return rand.Int()
	})

	for i:=0; i<128; i++ {
		set.Add(strconv.Itoa(i))
	}

	assert.Assert(t, set.Count() == 128, "Invalid set count: %d", set.Count())

	set.Clear()

	assert.Assert(t, set.Count() == 0, "Invalid set count: %d", set.Count())
}
