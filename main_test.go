package main

import (
	"errors"
	"gotest.tools/assert"
	"testing"
	"time"
)

func TestPromiseAll(t *testing.T) {

	result, e := PromiseAll(
		func() (interface{}, error) {
			time.Sleep(5 * time.Second)
			return nil, nil
		},
		func() (interface{}, error) {
			return 1+2, nil
		},
		func() (interface{}, error) {
			return 6*6, nil
		},
		func() (interface{}, error) {
			time.Sleep(1 * time.Second)
			return true, nil
		},
		func() (interface{}, error) {
			return "result", nil
		},
	)

	assert.NilError(t, e, "Error should be nil.")
	assert.Assert(t, len(result) == 5, "Invalid size of the result array.")
	assert.Assert(t, result[0] == nil, "Invalid output.")
	assert.Assert(t, result[1] == 3, "Invalid output.")
	assert.Assert(t, result[2] == 36, "Invalid output.")
	assert.Assert(t, result[3] == true, "Invalid output.")
	assert.Assert(t, result[4] == "result", "Invalid output.")
}

func TestPromiseAllError(t *testing.T) {

	result, e := PromiseAll(
		func() (interface{}, error) {
			return 1+2, nil
		},
		func() (interface{}, error) {
			return 6*6, nil
		},
		func() (interface{}, error) {
			time.Sleep(1 * time.Second)
			return true, nil
		},
		func() (interface{}, error) {
			return nil, errors.New("This is error.")
		},
	)

	assert.Assert(t, e  != nil, "Error should not be nil.")
	assert.Assert(t, e.Error()  == "This is error.", "Invalid error.")
	assert.Assert(t, len(result) == 0, "Result array must be empty.")
}
